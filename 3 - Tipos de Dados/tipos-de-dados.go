package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = 10000000000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// alias
	// INT32 =RUNE
	var numero3 rune = 12345
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	// Números reais
	var numeroReal1 float32 = 123.4
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123456.8
	fmt.Println(numeroReal2)

	numeroReal3 := 123456.8
	fmt.Println(numeroReal3)

	// string
	var str1 string = "Texto"
	fmt.Println(str1)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	texto := 5
	fmt.Println(texto)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno!")
	fmt.Println(erro)
}
