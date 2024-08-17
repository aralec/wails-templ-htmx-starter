package handlers

import (
	"encoding/json"
	"net/http"
	"wails-templ-hmtx-project/src/external/views"
	"wails-templ-hmtx-project/src/internal/core/ports"
	"wails-templ-hmtx-project/src/internal/core/services"
)

type CounterHandler struct {
	service ports.Counter
}

func NewCounterHandler() *CounterHandler {
	return &CounterHandler{
		service: services.NewCounterService(),
	}
}

func (ch *CounterHandler) ServeCounterView(w http.ResponseWriter, r *http.Request) {
	count := ch.service.GetCount()
	view := views.CounterView(count)
	view.Render(r.Context(), w)
}

func (ch *CounterHandler) Increment(w http.ResponseWriter, r *http.Request) {
	ch.service.Increment()
	count := ch.service.GetCount()
	json.NewEncoder(w).Encode(count)
}

func (ch *CounterHandler) Decrement(w http.ResponseWriter, r *http.Request) {
	ch.service.Decrement()
	count := ch.service.GetCount()
	json.NewEncoder(w).Encode(count)
}
