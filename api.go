package main

import (
	"encoding/json" // Para converter a struct em JSON e enviar como resposta HTTP. O pacote encoding/json é parte da biblioteca padrão do Go, então não precisamos de dependências externas.
	"net/http"      // Não precisa de um IIS ou de um servidor externo pesado para começar. O pacote nativo net/http já é um servidor de nível de produção.
)

// ListarMaoHandler seria como um método de um Controller (C#)
// ListarMaoHandler vai responder a mão do jogador em formato JSON
func ListarMaoHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Prepara os dados (Lógica já criada em cartas.go)
	baralho := NovoBaralho()
	baralho.Embaralhar()
	mao, _ := baralho.DarCartas(11)
	mao.Ordenar()

	// 2. Configura o cabeçalho para JSON
	w.Header().Set("Content-Type", "application/json")

	// 3. Transforma a struct em JSON e enviamos
	json.NewEncoder(w).Encode(mao)
}
