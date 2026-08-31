package main

import "fmt"

type Cliente struct {
	Nome string
	Idade int
	Endereco Endereco
	Email string
}

type Endereco struct {
	Rua string
	Numero int
	Cep string
	Estado string
}

func main() {
	cliente1 := Cliente{
		Nome: "test",
		Idade: 10,
		Endereco: Endereco{
			Rua: "teste",
			Numero: 123,
			Estado: "sp",
			Cep: "xxxx-xxx",
		},
		Email: "test@test",
	}

	fmt.Println(cliente1)

	cliente1.Endereco.Numero = 80

	fmt.Println(cliente1)
}
