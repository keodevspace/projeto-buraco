package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// --- ESTRUTURAS ---

// Carta representa uma única carta do jogo
type Carta struct {
	Valor  string
	Naipe  string
	Pontos int
}

// Baralho é um slice de cartas com métodos associados
type Baralho []Carta

// Mesa controla o que está exposto no jogo
type Mesa struct {
	Lixo  Baralho
	Monte Baralho
}

// --- MÉTODOS E FUNÇÕES ---

// NovoBaralho cria as 52 cartas padrão
func NovoBaralho() Baralho {
	baralho := Baralho{}
	naipes := []string{"Copas", "Ouros", "Paus", "Espadas"}
	valores := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, naipe := range naipes {
		for _, valor := range valores {
			pontuacao := 5 // 3, 4, 5, 6, 7

			if valor == "A" {
				pontuacao = 15
			} else if valor == "2" || valor == "8" || valor == "9" || valor == "10" || valor == "J" || valor == "Q" || valor == "K" {
				pontuacao = 10
			}

			novaCarta := Carta{Valor: valor, Naipe: naipe, Pontos: pontuacao}
			baralho = append(baralho, novaCarta)
		}
	}
	return baralho
}

// Embaralhar usa uma fonte de tempo para garantir aleatoriedade
func (b Baralho) Embaralhar() {
	fonte := rand.NewSource(time.Now().UnixNano())
	r := rand.New(fonte)

	for i := range b {
		posicaoAleatoria := r.Intn(len(b))
		b[i], b[posicaoAleatoria] = b[posicaoAleatoria], b[i]
	}
}

// Ordenar organiza a mão do jogador (Naipe > Valor)
func (b Baralho) Ordenar() {
	sort.Slice(b, func(i, j int) bool {
		if b[i].Naipe != b[j].Naipe {
			return b[i].Naipe < b[j].Naipe
		}
		// Nota: Aqui a ordenação de valor é alfabética (A, J, K, Q).
		// Em um sistema real, usaríamos o peso numérico da carta.
		return b[i].Valor < b[j].Valor
	})
}

// DarCartas retira X cartas do topo do baralho
func (b Baralho) DarCartas(quantidade int) (Baralho, Baralho) {
	return b[:quantidade], b[quantidade:]
}

// DescartarNoLixo adiciona uma carta ao lixo da mesa
func (m *Mesa) DescartarNoLixo(c Carta) {
	m.Lixo = append(m.Lixo, c)
}

func (m *Mesa) ComprarDoLixo() (Baralho, error) {
	if len(m.Lixo) == 0 {
		return nil, errors.New("o lixo está vazio")
	}
	cartas := m.Lixo
	m.Lixo = Baralho{} // Limpa o lixo
	return cartas, nil
}

// --- EXECUÇÃO ---

func main() {
	// 1. Setup inicial
	meuBaralho := NovoBaralho()
	meuBaralho.Embaralhar()
	mesa := Mesa{Lixo: Baralho{}}

	// 2. Distribuição
	maoJogador, resto := meuBaralho.DarCartas(11)
	mesa.Monte = resto

	// 3. Organização
	fmt.Println("Sua mão antes de organizar:")
	imprimirMao(maoJogador)

	maoJogador.Ordenar()

	fmt.Println("\nSua mão organizada:")
	imprimirMao(maoJogador)

	// 4. Simulação de Jogada
	fmt.Println("\n--- Rodada 1 ---")
	
	// Jogador descarta a última carta
	cartaDescarte := maoJogador[len(maoJogador)-1]
	maoJogador = maoJogador[:len(maoJogador)-1] // Remove da mão
	mesa.DescartarNoLixo(cartaDescarte)

	fmt.Printf("Você descartou [%s de %s]. Cartas no lixo: %d\n", 
		cartaDescarte.Valor, cartaDescarte.Naipe, len(mesa.Lixo))

	// 5. Somando pontos da mão
	total := 0
	for _, c := range maoJogador {
		total += c.Pontos
	}
	fmt.Printf("Pontos totais na mão: %d\n", total)
}

// Função auxiliar para imprimir as cartas de forma bonitinha
func imprimirMao(b Baralho) {
	for _, c := range b {
		fmt.Printf("[%s%s] ", c.Valor, string(c.Naipe[0])) // Mostra valor e a inicial do naipe
	}
	fmt.Println()
}