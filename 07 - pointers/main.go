package main

import "fmt"

func main() {
	var v1 int = 10
	var v2 int = v1 // v2 será uma cópia de v1

	fmt.Println(v1, v2)

	v1++
	fmt.Println(v1, v2)

	// Ponteiro é uma referência de memória
	var v3 int
	var pointer *int

	fmt.Println(v3, pointer)

	v3 = 100
	pointer = &v3

	// vai mostrar o valor de v3 e o endereço de memória que v3 está armazenado no computador
	fmt.Println(v3, pointer)

	// com * na frente, vai mostrar o valor armazenado, que é igual à v3
	fmt.Println(*pointer)

	// mudando o valor da variárvel o endereço de memória continua o mesmo, mas seu valor armazenado muda
	v3 = 150
	fmt.Println(v3, pointer)

	pointer = &v3
	fmt.Println(*pointer)
}
