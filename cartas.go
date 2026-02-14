package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Carta struct {
	Valor  string
	Naipe  string
	Pontos int
}

type Baralho []Carta

func NovoBaralho() Baralho {
	baralho := Baralho{}
	naipes := []string{"Copas", "Ouros", "Paus", "Espadas"}
	valores := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, naipe := range naipes {
		for _, valor := range valores {
			pontuacao := 5
			if valor == "A" {
				pontuacao = 15
			} else if valor == "2" || valor == "8" || valor == "9" || valor == "10" || valor == "J" || valor == "Q" || valor == "K" {
				pontuacao = 10
			}
			baralho = append(baralho, Carta{Valor: valor, Naipe: naipe, Pontos: pontuacao})
		}
	}
	return baralho
}

func (b Baralho) Embaralhar() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		posicaoAleatoria := r.Intn(len(b))
		b[i], b[posicaoAleatoria] = b[posicaoAleatoria], b[i]
	}
}

func (b Baralho) Ordenar() {
	sort.Slice(b, func(i, j int) bool {
		if b[i].Naipe != b[j].Naipe {
			return b[i].Naipe < b[j].Naipe
		}
		return b[i].Valor < b[j].Valor
	})
}

func (b Baralho) DarCartas(quantidade int) (Baralho, Baralho) {
	return b[:quantidade], b[quantidade:]
}

func ImprimirMao(b Baralho) {
	for _, c := range b {
		fmt.Printf("[%s%s] ", c.Valor, string(c.Naipe[0]))
	}
	fmt.Println()
}
