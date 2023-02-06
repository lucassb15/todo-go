package rotas

import (
	"net/http"

	"teste.com/pacotes/src/controllers"
)

var rotasIdade = []Rota{
	{
		URI:    "/idade",
		Metodo: http.MethodGet,
		Funcao: controllers.CarregarTelaDeIdade, // a Função que executa esta rota esta dentro do pacote "controllers"
	},
}
