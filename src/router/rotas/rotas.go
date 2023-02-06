package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da Aplicação Web
type Rota struct {
	URI    string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
}

// "Configurar" coloca todas as "rotas" dentro do "router"
func Configurar(router *mux.Router) *mux.Router {
	// Recebe um Router sem nenhuma rota dentro, configura todas as rotas e devolve o Router pronto.

	rotas := rotasIdade          // pega o slice de rotas de "idade.go"
	for _, rota := range rotas { // loop no alcance de []Rota que foi importado de "idade.go"
		router.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo) // passa as rotas
	}

	rotastodo := rotasTodo
	for _, rota := range rotastodo {
		router.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router // e retorna com as rotas configuradas
}
