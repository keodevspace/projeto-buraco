package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a rota
	http.HandleFunc("/minha-mao", ListarMaoHandler)

	// Rota nova para Descartar
	http.HandleFunc("/descartar", DescartarCartaHandler)

	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)

	// Sobe o servidor (Listen and Serve)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao subir servidor:", err)
	}
}
