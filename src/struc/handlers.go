package struc

import (
	"html/template"
	"net/http"
	"strings"
)

func HandleIndex(files []string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f := append(files, "templates/landing-page.html")
		tmpl := template.Must(template.ParseFiles(f...))
		tmpl.Execute(w, nil)
	})
}

func HandleForm(files []string) {
	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			f := append(files, "templates/1st-form.html")
			tmpl := template.Must(template.ParseFiles(f...))
			tmpl.Execute(w, nil)
			return
		}

		// Appel d'un personnage
		listPerso := FetchCharacter(r.FormValue("perso"))

		f := append(files, "templates/form.html")
		tmpl := template.Must(template.ParseFiles(f...))
		tmpl.Execute(w, listPerso)
	})
}

func HandlePerso(files []string) {
	http.HandleFunc("/perso", func(w http.ResponseWriter, r *http.Request) {
		idPerso := strings.TrimPrefix(r.URL.RequestURI(), "/perso?id=")
		chara := FetchCharacterByID(idPerso)

		f := append(files, "templates/fiche-perso.html")
		tmpl := template.Must(template.ParseFiles(f...))
		tmpl.Execute(w, chara)
		return
	})
}
