package main

import (
	"html/template"
	"log"
	"net/http"

	chat "github.com/okhrimko/websocketchat/server"
	shared "github.com/okhrimko/websocketchat/tools"
)

var index = template.Must(template.ParseFiles("./assets/html/index.html"))

func home(w http.ResponseWriter, r *http.Request) {
	index.Execute(w, r)
}

func main() {
	go chat.DefaultHub.Start()

	http.HandleFunc("/", home)
	http.HandleFunc("/ws", chat.WSHandler)

	port := shared.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
