package main

// VerificarCanastra recebe um grupo de cartas baixadas e valida
func VerificarCanastra(sequencia Baralho) string {
	if len(sequencia) < 7 {
		return "Não é canastra (menos de 7 cartas)"
	}

	// Por enquanto, vamos apenas dizer que se tem 7, é canastra
	return "CANASTRA! 🏆"
}
