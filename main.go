package main

import (
	"e2e-demo/app"
	"html/template"
	"log"
	"net/http"
)

func main() {
	store := &app.Store{}

	tmpl := template.Must(template.ParseFiles("web/index.html"))

	http.HandleFunc("/", app.HomeHandler(store, tmpl))

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
