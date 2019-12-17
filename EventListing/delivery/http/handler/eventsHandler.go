package handler

import (
	"html/template"
	"net/http"

	"github.com/birukbelay/Aprojects/eventListing/events"
)

type EventHandler struct {
	templ     *template.Template
	eventServ events.EventService
}

func NewEventHandler(T *template.Template, ES events.EventService) *EventHandler {
	return &EventHandler{templ: T, eventServ: ES}

}

func (ehd *EventHandler) Index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	events, err := ehd.eventServ.Events()
	if err != nil {
		panic(err)
	}
	ehd.templ.ExecuteTemplate(w, "grid.html", events)
}
