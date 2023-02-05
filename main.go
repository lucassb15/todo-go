package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

var tmpl *template.Template

type Todo struct {
	Item string
	Done bool
}

type PageData struct {
	Title string
	Todos []Todo
}

func todo(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseGlob("src/view/templates/index.html"))
	data := PageData{
		Title: "TODO List",
		Todos: []Todo{
			{Item: "Teste 1", Done: true},
			{Item: "Teste 2", Done: false},
			{Item: "Teste 3", Done: false},
			{Item: r.PostFormValue("inputUser"), Done: false},
		},
	}

	tmpl.Execute(w, data)
}

func idade(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseGlob("src/view/templates/idade.html"))
	tmpl.ExecuteTemplate(w, "idade.html", nil)
}

func main() {
	port := os.Getenv("PORT_WEB") // env port
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./src/view/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	mux.HandleFunc("/", todo)
	mux.HandleFunc("/idade", idade)
	fmt.Println("Server started at port:", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
