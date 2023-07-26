package forum

import (
	"net/http"
	"text/template"
)

func About(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("forum/static/about.html"))
	tmpl.Execute(w, Send{User: GetUserByCookies(w, r)})
}