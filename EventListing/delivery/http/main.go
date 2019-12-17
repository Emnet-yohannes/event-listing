package main

import (
	"database/sql"
	"html/template"
	"net/http"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"github.com/birukbelay/Aprojects/eventListing/delivery/http/handler"
	"github.com/birukbelay/Aprojects/eventListing/events/repository"
	"github.com/birukbelay/Aprojects/eventListing/events/services"
)

// some main comments
// some main edits branch changed
func main() {

	dbconn, err := sql.Open("postgres", "postgres://postgres:bura@localhost/eventlisting?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	if err := dbconn.Ping(); err != nil {
		panic(err)
	}

	templ := template.Must(template.ParseGlob("../../ui/templates/components/*.html"))

	EventRepo := repository.NewEventRepoImp(dbconn) // is a pointer to a struct with db conn obj as a field
	EventServ := services.NewEventSerImp(EventRepo) // is a pointer to a struct with the (--interface-- CategoryService  as a field

	eventsHandler := handler.NewEventHandler(templ, EventServ)

	fs := http.FileServer(http.Dir("../../ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", eventsHandler.Index)
	fmt.Println("server starting..")
	log.Fatal(http.ListenAndServe(":8181", nil))

}
