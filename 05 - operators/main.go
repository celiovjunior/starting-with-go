package main

import "fmt"

func main() {
	// -> OPERADORES ARITIMÉTICOS <-

	// mult := 2 * 4
	// fmt.Println(mult)

	// resto := 10 % 3
	// fmt.Println(resto)

	// var n1 int32 = 10
	// var n2 int32 = 4

	// soma := n1 + n2

	// soma := n1 + n2
	// fmt.Println(soma)

	// -> OPERADORES DE ATRIBUIÇÃO <-
	// var v1 string = "atribuindo"

	// v2 := "atribuindo de novo"

	// fmt.Println(v1, v2)

	// -> OPERADORES RELACIONAIS <-
	// fmt.Println(1 > 2)
	// fmt.Println(1 >= 2)
	// fmt.Println(1 != 2)

	// -> OPERADORES LÓGICOS <-
	// v, f := true, false
	// fmt.Println(v && v)
	// fmt.Println(v || f)

	// -> OPERADOR "TERNÁRIO"
	var confirm bool = true
	var msg string
	if !confirm {
		msg = "email não enviado"
		fmt.Println(msg)
	} else {
		msg = "email enviado"
		fmt.Println(msg)
	}

}
