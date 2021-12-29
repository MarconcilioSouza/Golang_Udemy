package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1"
	fmt.Println(variavel1)

	variavel2 := "Variável 2"

	fmt.Println(variavel2)

	var (
		variavel3 string = "Lapís"
		variavel4 string = "Papel"
	)

	fmt.Println(variavel3)
	fmt.Println(variavel4)

	variavel5, variavel6 := "var 5", "var 6"

	fmt.Println(variavel5)
	fmt.Println(variavel6)
}
