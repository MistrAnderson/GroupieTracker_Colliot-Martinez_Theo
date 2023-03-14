package main

import (
	"Groupie/struc"
	"fmt"
	"html/template"
	"net/http"
)

const port = ":8080"

func main() {

	files := []string{
		"templates/index.html",
		"templates/fiche-perso.html",
		"templates/header.html",
		"templates/form.html"}

	tmpl := template.Must(template.ParseFiles(files...))

	// Ajout des fichiers statiques
	static := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", static))

	fmt.Println("(http://localhost:8080)", "Server started on port ", port)

	// Gère la route "/"
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		tmpl.Execute(w, nil)
	})

	// Gère la route "/form"
	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
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
