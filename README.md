# Choose your own adventure style web app

To start run from the root of the project, the app should start on port 3000:

```bash
go run .
```

There are some options to customise starting the app:

- port (the port to start the app on)

```bash
go run . -port 4567
```

- file (the story in json format to start the app with), the story should be in the src/stories folder and should have the correct format (see src/stories/story.json as an example)

```bash
go run . -file another_story.json
```
