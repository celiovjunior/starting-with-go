package main

import "fmt"

func main() {
	// -> Estrutura tradicional
	// idade := 12

	// if idade >= 18 {
	// 	fmt.Println("pode dirigir")
	// } else {
	// 	fmt.Println("não pode")
	// }

	// -> Estrutura com if init
	// if init é a capacidade de criar uma var nova dentro da condicional do if
	mediaGeral := 4.8

	// mediaAluno não foi declarada antes, mas foi agora no if init
	if mediaAluno := mediaGeral; mediaAluno <= 6.9 {
		fmt.Println("aluno de recuperação")
	} else {
		fmt.Println("aluno aprovado")
	}

	alturaUsuario := 1.5
	pesoUsuario := 85.0

	imcUsuario := pesoUsuario / (alturaUsuario * alturaUsuario)

	if imc := imcUsuario; imc <= 18 {
		fmt.Println("abaixo do peso")
	} else if imc > 18 && imc <= 26 {
		fmt.Println("peso ideal")
	} else if imc > 26 {
		fmt.Println("acima do peso")
	}
}
