package main

import (
	"fmt"
	"groupie/src/struc"
	"html/template"
	"net/http"
)

const port = ":8080"

func main() {

	files := []string{
		"templates/index.html",
		"templates/fiche-perso.html"}

	tmpl := template.Must(template.ParseFiles(files...))

	// Ajout du CSS
	fs := http.FileServer(http.Dir("style"))
	http.Handle("/style/", http.StripPrefix("/style/", fs))

	fmt.Println("(http://localhost:8080)", "Server started on port ", port)

	// Gère la route "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		// Appel d'un personnage
		c1 := struc.FetchCharacter(r.FormValue("perso"))

		tmpl.Execute(w, c1)
	})

	http.ListenAndServe(":8080", nil)
}
