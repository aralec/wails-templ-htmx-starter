package main

import (
	"net/http"
	"wails-templ-htmx-project/src/cmd/render"
)

// main runs the core application (server-side code API)
func main() {
	r := render.GetRenderingMux() // TODO : main should only serve the core backend API. New adapters must be added to the repository.
	http.ListenAndServe(":8080", r)
}
