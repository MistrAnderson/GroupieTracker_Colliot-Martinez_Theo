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
		"templates/header.html"}

	// Ajout des fichiers statiques
	static := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", static))

	fmt.Println("(http://localhost:8080)", "Server started on port ", port)

	// Gère la route "/"
	struc.HandleIndex(files)

	// Gère la route "/formPerso"
	struc.HandleFormPerso(files)

	//Gère la route "/formPerso/perso"
	struc.HandlePerso(files)

	http.ListenAndServe(":8080", nil)
}
