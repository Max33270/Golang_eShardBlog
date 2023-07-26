package forum

import (
	"net/http"
	"text/template"
)

func Summary(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("forum/static/summary.html"))
	tmpl.Execute(w, Send{User: GetUserByCookies(w, r)})
}