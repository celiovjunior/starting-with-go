package main

import "fmt"

func main() {
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println(i)
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println(i)

	// for j := 0; j < 10; j += 5 {
	// 	fmt.Println(j)
	// 	time.Sleep(time.Second)
	// }

	// nomes := [3]string{"rodrigo", "miguel", "caio"}

	// for index, nome := range nomes {
	// 	fmt.Println(index, nome)
	// }

	// for index, letra := range "PALAVRA" {
	// 	fmt.Println(index, string(letra))
	// }

	// usuario := map[string]string{
	// 	"nome":      "celio",
	// 	"sobrenome": "viana",
	// }

	// fmt.Println(usuario)

	// for chave, valor := range usuario {
	// 	fmt.Println(chave, valor)
	// }

	type usuario struct {
		nome      string
		sobrenome string
		cidade    string
		anoNasc   int16
	}

	u2 := usuario{"celio", "viana", "fortaleza", 1994}

	// NÃO funciona em struct
	// for chave, valor := range u2 {
	// 	fmt.Println(chave, valor)
	// }

	fmt.Println(u2)

	// fmt.Println(nomes)

	for {
		fmt.Println(1)
	}
}
