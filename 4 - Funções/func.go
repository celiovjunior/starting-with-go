package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// As funções no golang permitem que você tenha mais de um retorno
func mathCalcs(n1, n2 int32) (int32, int32) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	// executando a função declarada fora do main
	soma := somar(28, 25)
	fmt.Println(soma)

	// construindo uma variável do tipo func
	var f = func(txt string) {
		fmt.Println(txt)
	}

	// chamando a var do tipo func
	f("executando f com paramentro")

	var f2 = func(txt string) string {
		return txt
	}

	resultado := f2("executou f2")
	fmt.Println(resultado)

	// usando a func com 2 retornos
	// resSoma, resSub := mathCalcs(10, 15)
	// fmt.Println(resSoma, resSub)

	// ignorando um dos retornos
	_, resSub := mathCalcs(2022, 1994)
	fmt.Println(resSub)

}
