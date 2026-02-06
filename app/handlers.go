package app

import (
	"html/template"
	"net/http"
)

func HomeHandler(store *Store, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			item := r.FormValue("item")
			if item != "" {
				store.Add(item)
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		_ = tmpl.Execute(w, store.List())
	}
}
