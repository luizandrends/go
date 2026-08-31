package main

import (
	"errors"
	"fmt"
)

type User struct {
	nome string
	idade int32
	peso float64
	email string
}

type Carros struct {
	id int32
	nome string
	cor string
	preco float64
	marca string
}

func main () {
	nums := []int{1, 2, 3, 4, 5}

	for key, value := range nums {
		fmt.Println(key, value)
	}

	users := map[string]User{
		"Luiz": {nome: "Luiz André", idade: 28, peso: 90.0, email: "teste@teste.com"},
	}

	carros := []Carros {
		{id: 1, nome: "Golf", cor: "prata", preco: 79900.0, marca: "Volkswagen"},
		{id: 2, nome: "Macan", cor: "preto", preco: 350000.0, marca: "Porsche"},
		{id: 3, nome: "M3", cor: "laranja", preco: 450000.0, marca: "BMW"},
		{id: 4, nome: "GLC 63", cor: "preto", preco: 550000.0, marca: "Mercedes"},
		{id: 5, nome: "RS3", cor: "preto", preco: 450000.0, marca: "Audi"},
	}

	for k, v := range users {
		fmt.Println(k, v)
	}

	for _, v := range carros {
		if v.preco > 100000.0 && v.cor == "laranja" && v.marca == "Mercedes" {
			fmt.Println(v.nome)
		} else {
			notFoundErr()
		}
	}
}

func notFoundErr() error {
	return errors.New("Carro não encontrado")
}