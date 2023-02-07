package main

import (
	"teste.com/pacotes/src/config/database/mongodb"
)

func main() {

	// utils.CarregarTemplates()
	// r := router.Gerar()
	// port := os.Getenv("PORT_WEB") // env port
	// fmt.Println("Servidor iniciado", port)
	// log.Fatal(http.ListenAndServe(port, r))
	mongodb.NewMongoDBConnection()
}
