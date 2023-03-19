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

//		PARTIE PERSO

func HandleFormPerso(files []string) {
	http.HandleFunc("/formPerso", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			f := append(files, "templates/Perso/1st-form-perso.html")
			tmpl := template.Must(template.ParseFiles(f...))
			tmpl.Execute(w, nil)
			return
		}

		// Appel d'un personnage
		listPerso := FetchCharacter(r.FormValue("perso"))

		to_parse := []string{"templates/Perso/form-perso.html", "templates/Perso/liste-perso.html"}
		f := append(files, to_parse...)
		tmpl := template.Must(template.ParseFiles(f...))
		tmpl.Execute(w, listPerso)
	})
}

func HandlePerso(files []string) {
	http.HandleFunc("/Perso", func(w http.ResponseWriter, r *http.Request) {
		idPerso := strings.TrimPrefix(r.URL.RequestURI(), "/Perso?id=")
		chara := FetchCharacterByID(idPerso)

		f := append(files, "templates/Perso/fiche-perso.html")
		tmpl := template.Must(template.ParseFiles(f...))
		tmpl.Execute(w, chara)
		return
	})
}

//		PARTIE CREATEURS

func HandleFormCreator(files []string) {
	http.HandleFunc("/formCreators", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			f := append(files, "templates/Creators/1st-form-creators.html")
			tmpl := template.Must(template.ParseFiles(f...))
			tmpl.Execute(w, nil)
			return
		}

		// Appel d'un créateur
		listPerso := FetchCreator(r.FormValue("perso"))

		f := append(files, "templates/Creators/form-creators.html")
		tmpl := template.Must(template.ParseFiles(f...))
		tmpl.Execute(w, listPerso)
	})
}
