package rotas

import (
	"net/http"

	"teste.com/pacotes/src/controllers"
)

var rotasTodo = []Rota{
	{
		URI:    "/index",
		Metodo: http.MethodGet,
		Funcao: controllers.CarregarTelaDeIndex,
	},
	{
		URI:    "/",
		Metodo: http.MethodGet,
		Funcao: controllers.CarregarTelaDeIndex,
	},
}
