package struc

import (
	"html/template"
	"net/http"
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
		c1 := FetchCharacter(r.FormValue("perso"))

		f := append(files, "templates/form.html")
		tmpl := template.Must(template.ParseFiles(f...))
		tmpl.Execute(w, c1)
	})
}
