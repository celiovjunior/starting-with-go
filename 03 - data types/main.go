package main

import (
	"errors"
	"fmt"
)

func main() {

	// -> INTEIROS <-

	// var numero int = 10000000000
	// numero2 := 265465425624626

	// var numero3 uint = 10000

	// fmt.Println(numero2)
	// fmt.Println(numero)
	// fmt.Println(numero3)

	// alias
	// int32 == rune
	// byte == uint8

	// var numero4 rune = 54654
	// fmt.Println(numero4)

	// -> REAIS <-

	// var realn float32 = 123.45
	// fmt.Println(realn)

	// var realn2 float64 = 123.486534674646464564654
	// fmt.Println(realn2)

	// -> STRINGS <-

	// meuNome := "celio"
	// fmt.Println(meuNome)

	// var minhaCidade string = "fortaleza"
	// fmt.Println(minhaCidade)

	// char := 'B'
	// fmt.Println(char)

	// var text int16
	// fmt.Println(text)

	var entrar bool = true
	fmt.Println(entrar)

	// -> ERROR <-
	var errou error = errors.New("A aplicação não funcionou de forma esparada")
	fmt.Println(errou)

}
