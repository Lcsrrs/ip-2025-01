package main

import (
	"net/http"
	"text/template"
)

var temp *template.Template

func main() {

	fs := http.FileServer(http.Dir("/pagina_web/static"))
	temp, _ = template.ParseGlob("static/*.html")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/sobre", sobreHandler)
	http.HandleFunc("/projetos", projetosHandler)
	http.HandleFunc("/contato", contatoHandler)
	http.Handle("./static", http.StripPrefix("./static", fs))

	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "index.html", nil)
}

func sobreHandler(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "sobre.html", nil)
}

func projetosHandler(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "projetos.html", nil)
}

func contatoHandler(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "contato.html", nil)
}
