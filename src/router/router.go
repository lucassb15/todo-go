package router

import (
	"github.com/gorilla/mux"
	"teste.com/pacotes/src/router/rotas"
)

// retorna um router com todas as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
