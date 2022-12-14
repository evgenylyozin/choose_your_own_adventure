package adventure

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"path"
	"strings"
)

func init() {
	tmpl = template.Must(template.ParseFiles(path.Join("./", "src", "templates", "template.html")))
}

var tmpl *template.Template

// HandlerOption ...
type HandlerOption func(h *handler)

// WithTemplate ...
func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

// NewHandler ...
func NewHandler(s Story, opts ...HandlerOption) http.Handler {
	h := handler{s, tmpl}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

type handler struct {
	s Story
	t *template.Template
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urlPath := strings.TrimSpace(r.URL.Path)

	if urlPath == "" || urlPath == "/" {
		urlPath = "/intro"
	}

	urlPath = urlPath[1:]

	if chapter, ok := h.s[urlPath]; ok {
		err := h.t.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "Chapter not found", http.StatusNotFound)
	return
}

// JSONStory ...
func JSONStory(r io.Reader) (Story, error) {

	d := json.NewDecoder(r)

	var story Story

	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

// Story ...
type Story map[string]Chapter

// Chapter ...
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option ...
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
