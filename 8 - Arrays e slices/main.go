package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays e slices")

	// var array1 [5]string
	// array1[0] = "index 1"
	// fmt.Println(array1)

	// array2 := [5]int{1, 2, 3, 4, 5}
	// fmt.Println(array2)

	// ele não permite adicionar outro elemento dessa forma
	// array2[5] = 6

	// array3 := [...]int{1, 2, 3, 4, 5}
	/*fmt.Println(array3)*/ // apesar de não especificar seu tamanho, após a declaração ele fica com tamanho fixo

	// -> SLICE <-

	// slice := []int{1, 2, 3, 4, 5, 8, 6, 4, 8, 2, 5, 6}

	// fmt.Println(reflect.TypeOf(slice))
	// fmt.Println(reflect.TypeOf(array3))

	// slice = append(slice, 18)
	// fmt.Println(slice)

	// slice2 := array2[1:3]
	// fmt.Println(slice2)

	// array2[1] = 25
	// fmt.Println(slice2)

	// -> ARRAYS INTERNOS <-
	slice := make([]string, 3, 5)
	fmt.Println(slice)

	slice = append(slice, "oi")
	fmt.Println(slice)

	slice = append(slice, "tudo")

	slice = append(slice, "bem")
	fmt.Println(len(slice)) // length
	fmt.Println(cap(slice)) // capacity

	slice2 := make([]float32, 5)
	fmt.Println(slice2)

	slice2 = append(slice2, 0.4)
	fmt.Println(cap(slice2))

}
