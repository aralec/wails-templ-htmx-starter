package main

import (
	"net/http"
	"wails-templ-htmx-project/src/cmd/render"
)

// main runs the core application (server-side code + templating)
func main() {
	r := render.GetRenderingMux()
	http.ListenAndServe(":8080", r)
}
