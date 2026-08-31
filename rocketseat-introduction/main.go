package main

import (
	"fmt"
)

func main() {
	// meet.SayHello()
	// meet.Say("teste123")

	// var texto string = "ola"

	// var num = 22
	// num += 5

	// var a string
	// var b bool
	// c := "teste"

	// const test string = "my-const"
	// const d = "my-const"

	// var idade int = 30
	// var contador int32 = 2
	// var indice int8 = 1

	// var floatNumber float32 = 1.1
	// var pi float64 = 3.14
	// var raio float64 = 2.5
	// var area = pi * raio * raio

	// var maior bool = 10 > 5
	// var menor bool = 10 < 5

	// var hello string = "Olá, mundo!"
	// var question string = " Como vai?"

	// var meet = hello + question

	// fmt.Println(strings.ToUpper(meet))
	// fmt.Println(strings.Contains(meet, "mundo"))

	// var gavetas [2]string
	// gavetas[0] = "copos"
	// gavetas[1] = "panos"
	// fmt.Println(gavetas[1:2])
	// //slice[x:x-1]

	// var guardaroupas []string
	// guardaroupas = append(guardaroupas, "camisa", "blusa")
	// fmt.Println(len(guardaroupas))

	var pessoas = map[string]int{}
	var num int
	fmt.Println(num)
	pessoas["Luiz"] = 28
	pessoas["Leo"] = 32
	fmt.Println(pessoas)

	if idade, ok := pessoas["Luiz"]; ok {
		fmt.Println("Pessoa existe no map", idade, ok)
	} else {
		fmt.Println("Pessoa não existe no map")
	}

	delete(pessoas, "Leo")
	fmt.Println(pessoas)
}
