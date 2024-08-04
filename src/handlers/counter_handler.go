package handlers

import (
	"net/http"
	"wails-templ-hmtx-project/src/services"
	"wails-templ-hmtx-project/src/views"
)

type CounterHandler struct {
	service *services.CounterService
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
	ch.ServeCounterView(w, r)
}

func (ch *CounterHandler) Decrement(w http.ResponseWriter, r *http.Request) {
	ch.service.Decrement()
	ch.ServeCounterView(w, r)
}