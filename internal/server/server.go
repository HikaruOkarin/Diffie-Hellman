package server

import (
	"html/template"
	"net/http"
	"strconv"
)

type Data struct {
	Output   string
	Prime    string
	G        string
	Secret_a string
	Secret_b string
	A, B     string
	K_a, K_b string
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
	data.G = r.FormValue("generator")
	data.Secret_a = r.FormValue("key_a")
	data.Secret_b = r.FormValue("key_b")
	p, _ := strconv.Atoi(data.Prime)
	g, _ := strconv.Atoi(data.G)
	a, _ := strconv.Atoi(data.Secret_a)
	b, _ := strconv.Atoi(data.Secret_b)
	sec_A := modExp(g, a, p)
	sec_B := modExp(g, b, p)
	data.A = strconv.Itoa(sec_A)
	data.B = strconv.Itoa(sec_B)

	k_a := modExp(sec_B, a, p)
	k_b := modExp(sec_A, b, p)
	data.K_a = strconv.Itoa(k_a)
	data.K_b = strconv.Itoa(k_b)

	_ = tmpl.Execute(w, data)
}

func modExp(base, exponent, modulus int) int {
	result := 1
	base = base % modulus
	for exponent > 0 {
		if exponent%2 == 1 {
			result = (result * base) % modulus
		}
		exponent = exponent >> 1
		base = (base * base) % modulus
	}
	return result
}
