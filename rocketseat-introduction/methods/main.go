package main

import "fmt"

type Pessoa struct {
	Nome string
	Idade int
}

func (p Pessoa) Apresentar() {
	fmt.Printf("Olá meu nome é %s e tenho %d anos.", p.Nome, p.Idade)
}

func main() {
	p1 := Pessoa{Nome: "Luiz", Idade: 29}
	p1.Apresentar()
}
