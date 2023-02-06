package utils

import (
	"html/template"
	"net/http"
)

//vai interagir com o pacote html/template do go

var Templates *template.Template

// Insere os templates html na variável templates
func CarregarTemplates() {
	Templates = template.Must(template.ParseGlob("views/*.html")) // passa uma forma de identificar os arquivos que vai ser considerado templates fica na pasta "views"
}

// Renderiza uma página HTML na tela
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	Templates.ExecuteTemplate(w, template, dados) // os controllers utilizam essa função para executar os templates
}
