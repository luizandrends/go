package main

import (
	"errors"
	"fmt"
)

func main() {
	// nota := 96

	// if nota >= 90 {
	// 	fmt.Println("Aprovado com distinção")
	// } else if nota >= 70 {
	// 	fmt.Println("Aprovado")
	// } else {
	// 	fmt.Println("Aluno foi reprovado")
	// }

	if err := thisIsAnError(); err != nil {
		fmt.Println(err.Error())
	}

	err := thisIsAnError()
	fmt.Println(err)

	players := map[string]int{
		"lais": 26,
		"luiz": 30,
	}

	if value, ok := players["teste123"]; ok {
		fmt.Println("pontos:", value, ok)
	}
}

func thisIsAnError() error {
	return errors.New("Isto é um erro")
}