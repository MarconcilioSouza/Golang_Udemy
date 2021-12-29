package main

import "fmt"

func main() {
	// ARITIMETICOS
	soma := 1 + 2
	subtracao := 2 - 1
	divisao := 4 / 2
	multiplicacao := 2 * 2
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2

	fmt.Println(soma2)

	// ATRIBUIÇÃO

	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// OPERADORES LÓGICOS
	fmt.Println("------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//OPERADORES UNÁRIOS
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero -= 15
	numero *= 3
	numero /= 10
	numero %= 3

	fmt.Println(numero)

}
