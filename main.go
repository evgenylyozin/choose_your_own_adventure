package main

import (
	"adventure/adventure"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the app on")
	file := flag.String("file", "story.json", "the JSON file with the CYOA story")
	flag.Parse()

	f, err := os.Open(path.Join("./src/stories", *file))
	if err != nil {
		panic(err)
	}

	story, err := adventure.JSONStory(f)

	if err != nil {
		panic(err)
	}

	h := adventure.NewHandler(story)
	fmt.Printf("Starting the server on port, %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
