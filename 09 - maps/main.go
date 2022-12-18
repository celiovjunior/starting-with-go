package main

import "fmt"

func main() {
	user := map[string]string{
		"name":    "celio",
		"surname": "viana",
	}

	fmt.Println(user["surname"])

	user2 := map[string]map[string]string{
		"nome": {
			"primeiro": "celio",
			"segundo":  "viana",
		},

		"curso": {
			"nome": "adm",
		},
	}

	delete(user2, "nome")
	fmt.Println(user2)

}
