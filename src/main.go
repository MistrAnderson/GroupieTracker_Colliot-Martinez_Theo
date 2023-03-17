package main

import (
	"Groupie/struc"
	"fmt"
	"net/http"
)

const port = ":8080"

func main() {

	files := []string{
		"templates/index.html",
		"templates/liste-perso.html",
		"templates/header.html",
		"templates/form.html"}

	// Ajout des fichiers statiques
	static := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", static))

	fmt.Println("(http://localhost:8080)", "Server started on port ", port)

	// Gère la route "/"
	struc.HandleIndex(files)

	// Gère la route "/form"
	struc.HandleForm(files)

	//Gère la route "/perso"
	struc.HandlePerso(files)

	http.ListenAndServe(":8080", nil)
}
