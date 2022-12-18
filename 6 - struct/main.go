package main

import "fmt"

// Struct é o mais próximo de uma Class (do POO) no Golang.

// type user struct {
// 	name    string
// 	age     uint8
// 	address address
// }

// type address struct {
// 	street string
// 	number uint32
// }

type person struct {
	name string
	age  uint16
}

type student struct {
	person // student extends person
	course string
	city   string
}

func main() {
	fmt.Println("*** Struct ***")

	// var u1 user
	// u1.name = "Celio"
	// u1.age = 28
	// fmt.Println(u1)

	// u2 := user{"Rodrigo", 25}
	// fmt.Println(u2)

	// u3 := user{name: "Miguel"}
	// fmt.Println(u3)

	// ads1 := address{"Rua ali", 1478}

	// u1 := user{"Lucas", 19, ads1}
	// fmt.Println(u1)

	//-> O mais perto de Herança com Golang <-
	p1 := person{"celio", 28}
	s1 := student{p1, "ads", "fortaleza"}
	fmt.Println(s1.age)

}
