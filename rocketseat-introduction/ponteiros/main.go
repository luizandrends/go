package main

import "fmt"

type Pessoa struct {
	Nome string
}

func main() {
	var p1 Pessoa = Pessoa{Nome: "lais"}
	var p2 Pessoa = Pessoa{Nome: "luiz"}

	// Endereço de memória
	var p3 *Pessoa = &p1
	p3.Nome = "Angela"
	fmt.Println(&p1.Nome)
	fmt.Println(p1.Nome)
	fmt.Println(p2)
}