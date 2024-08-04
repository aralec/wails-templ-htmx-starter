package handlers

import (
	"net/http"
	"wails-templ-hmtx-project/src/views"
)

func HandleBar(w http.ResponseWriter, r *http.Request) {
	view := views.View2()
	view.Render(r.Context(), w)
}
