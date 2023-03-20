package struc

import (
	"html/template"
	"net/http"
	"strings"
)

func HandleIndex(files []string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			HandleNotFound(files, w, r)
			return
		}
		f := append(files, "templates/landing-page.html")
		tmpl := template.Must(template.ParseFiles(f...))
		tmpl.Execute(w, nil)
	})
}

//		PARTIE PERSO

func HandleFormPerso(files []string) {
	http.HandleFunc("/formPerso", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/formPerso" {
			HandleNotFound(files, w, r)
			return
		}
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
		if r.URL.Path != "/Perso" {
			HandleNotFound(files, w, r)
			return
		}
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
		if r.URL.Path != "/formCreators" {
			HandleNotFound(files, w, r)
			return
		}
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

//		PARTIE NOT-YET

func HandleNotYet(files []string) {
	http.HandleFunc("/NotYet", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/NotYet" {
			HandleNotFound(files, w, r)
			return
		}
		f := append(files, "templates/not-yet.html")
		tmpl := template.Must(template.ParseFiles(f...))
		tmpl.Execute(w, nil)
	})
}

// PARTIE 404
func HandleNotFound(files []string, w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(404)
	f := append(files, "templates/not-found.html")
	tmpl := template.Must(template.ParseFiles(f...))
	tmpl.Execute(w, nil)
}
