package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivos structs")

	var u usuario
	fmt.Println(u)
	u.nome = "Marconcilio"
	u.idade = 34
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua Prjetada ", 2}

	usuario2 := usuario{"Marco", 32, enderecoExemplo}
	fmt.Println(usuario2)
}
