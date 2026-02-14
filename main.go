package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a rota
	http.HandleFunc("/minha-mao", ListarMaoHandler)

	fmt.Println("Servidor de Buraco rodando em http://localhost:8080/minha-mao")

	// Sobe o servidor (Listen and Serve)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao subir servidor:", err)
	}
}
