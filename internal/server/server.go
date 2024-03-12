package server

import (
	"html/template"
	"net/http"
)

type Data struct {
	Output string
	Prime  string
}

func Runserver() {
	http.HandleFunc("/", home)
	http.HandleFunc("/diffie-hellman", Diffie)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	data := Data{Output: "", Prime: ""}
	tmpl, _ := template.ParseFiles("ui/index.html")

	_ = tmpl.Execute(w, data)
}

func Diffie(w http.ResponseWriter, r *http.Request) {
	data := Data{}
	tmpl, _ := template.ParseFiles("ui/index.html")

	data.Prime = r.FormValue("prime")

	_ = tmpl.Execute(w, data)
}
