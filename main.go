package main

import (
	"fmt"
	"groupie/src/struc"
	"html/template"
	"net/http"
)

const port = ":8080"

func main() {
	c1 := struc.FetchCharacter("Daredevil")

	fmt.Println("(http://localhost:8080)", "Server started on port ", port)

	// Gère la route "/"
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {

		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, c1)
	})

	http.ListenAndServe(":8080", nil)
}
