package forum

import (
	"net/http"
	"text/template"
)

func AboutCompany(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("forum/static/about_company.html"))
	tmpl.Execute(w, Send{User: GetUserByCookies(w, r)})
}