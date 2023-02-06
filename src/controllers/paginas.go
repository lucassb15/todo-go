package controllers

import (
	"net/http"

	"teste.com/pacotes/src/utils"
)

// Renderizar a tela Idade
func CarregarTelaDeIdade(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "idade.html", nil)
}

// Renderizar a tela Index
func CarregarTelaDeIndex(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "index.html", nil)
}
