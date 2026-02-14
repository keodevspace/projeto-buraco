package main

import "errors"

type Mesa struct {
	Lixo  Baralho
	Monte Baralho
}

func (m *Mesa) DescartarNoLixo(c Carta) {
	m.Lixo = append(m.Lixo, c)
}

func (m *Mesa) ComprarDoLixo() (Baralho, error) {
	if len(m.Lixo) == 0 {
		return nil, errors.New("o lixo está vazio")
	}
	cartas := m.Lixo
	m.Lixo = Baralho{}
	return cartas, nil
}
