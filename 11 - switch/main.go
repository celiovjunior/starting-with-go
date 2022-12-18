package main

import "fmt"

func diaDaSemana(numero int) string {
	var dia string
	switch {
	case numero == 1:
		dia = "Domingo"
		fallthrough // cláusula faz a variável cair na próxima condição
	case numero == 2:
		dia = "Segunda"
	case numero == 3:
		dia = "Terça"
	case numero == 4:
		dia = "Quarta"
	case numero == 5:
		dia = "Quinta"
	case numero == 6:
		dia = "Sexta"
	case numero == 7:
		dia = "Sábado"
	default:
		dia = "Número inválido"
	}

	return dia
}

// func diaDaSemana2(numero int) string {
// 	switch {
// 	case numero == 1:
// 		return "Domingo"
// 	default:
// 		return "nulo"
// 	}
// }

func main() {
	// vai receber o retorno da função
	// dia := diaDaSemana(2)
	// dia2 := diaDaSemana2(0)

	qualDia := diaDaSemana(1)

	// printanto o valor de 'qualDia'
	fmt.Println(qualDia)
}
