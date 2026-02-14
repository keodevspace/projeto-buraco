package main

import (
	"encoding/json" // Para converter a struct em JSON e enviar como resposta HTTP. O pacote encoding/json é parte da biblioteca padrão do Go, então não precisamos de dependências externas.
	"fmt"
	"net/http" // Não precisa de um IIS ou de um servidor externo pesado para começar. O pacote nativo net/http já é um servidor de nível de produção.
)

// ListarMaoHandler seria como um método de um Controller (C#)
// ListarMaoHandler vai responder a mão do jogador em formato JSON, ou seja, funciona como um GET
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

// PedidoDescarte representa o dado que o frontend (celular/web) vai enviar, ou sejam funciona como o POST
type PedidoDescarte struct {
	IndiceDaCarta int `json:"indice"`
}

func DescartarCartaHandler(resposta http.ResponseWriter, requisicao *http.Request) {
	// 1. Segurança: Só aceita se o método for POST
	if requisicao.Method != http.MethodPost {
		http.Error(resposta, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// 2. Lendo o JSON que o usuário enviou
	var pedido PedidoDescarte
	decodificador := json.NewDecoder(requisicao.Body)
	erro := decodificador.Decode(&pedido)

	if erro != nil {
		http.Error(resposta, "Erro ao ler o pedido", http.StatusBadRequest)
		return
	}

	// 3. Lógica de negócio (Simplificada para o teste)
	// Aqui, buscaria a mão do jogador real, mas é apenas simulacao
	fmt.Printf("O jogador quer descartar a carta da posição: %d\n", pedido.IndiceDaCarta)

	// 4. Responde que deu tudo certo
	resposta.WriteHeader(http.StatusAccepted)
	resposta.Write([]byte("Carta descartada com sucesso!"))
}
