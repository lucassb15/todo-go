package controllers

import (
	"net/http"

	"teste.com/pacotes/src/utils"
)

// Renderiza a tela de Idade
func CarregarTelaDeIdade(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "idade.html", nil)
}
func CarregarTelaDeIndex(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "index.html", nil)
}
