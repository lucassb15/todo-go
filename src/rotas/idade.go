package rotas

import (
	"net/http"

	"teste.com/pacotes/src/controllers"
)

var rotasIdade = []Rota{
	{
		URI:    "/idade",
		Metodo: http.MethodGet,
		Funcao: controllers.CarregarTelaDeIdade,
	},
}

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
