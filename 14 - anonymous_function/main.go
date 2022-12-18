package main

import "fmt"

func main() {
	returned := func(text string) string {
		return fmt.Sprintf("received -> %s, %d", text, 25)
	}("hello world parsed param")
	fmt.Println(returned)
}
