package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"teste.com/pacotes/src/router"
	"teste.com/pacotes/src/utils"
)

// var tmpl *template.Template

// type Todo struct {
// 	Item string
// 	Done bool
// }

// type PageData struct {
// 	Title string
// 	Todos []Todo
// }

// func todo(w http.ResponseWriter, r *http.Request) {
// 	tmpl = template.Must(template.ParseFiles("src/templates/index.html"))
// 	data := PageData{
// 		Title: "TODO List",
// 		Todos: []Todo{
// 			{Item: "Teste 1", Done: true},
// 			{Item: "Teste 2", Done: false},
// 			{Item: "Teste 3", Done: false},
// 			{Item: r.PostFormValue("inputUser"), Done: false},
// 		},
// 	}

// 	tmpl.Execute(w, data)
// }

// func idade(w http.ResponseWriter, r *http.Request) {
// 	tmpl = template.Must(template.ParseGlob("src/view/templates/idade.html"))
// 	tmpl.ExecuteTemplate(w, "idade.html", nil)
// }

func main() {
	// utils.CarregarTemplates()
	// r := router.Gerar()           // rotas
	// fmt.Println("Servidor iniciado :", port)
	// log.Fatal(http.ListenAndServe(port, r))
	utils.CarregarTemplates()
	r := router.Gerar()
	port := os.Getenv("PORT_WEB") // env port
	fmt.Println("Servidor iniciado :", port)
	log.Fatal(http.ListenAndServe(port, r))
}

// func main
// tmpl = template.Must(template.ParseGlob("views/templates/idade.html"))
// mux := http.NewServeMux() linha 55 foi retirado "mux" e colocado "r" log.Fatal(
// fileServer := http.FileServer(http.Dir("./views/static"))
// r.Handle("/static/", http.StripPrefix("/static/", fileServer)) // foi substituito "mux.Handle" para "r.Handle"
// r.HandleFunc("/", todo)                                        // foi substituito "mux.Handle" para "r.Handle"
// foi substituito "mux.Handle" para "r.Handle"
