package main

import "fmt"

func main() {
	resultado := soma(2, 2)

	fmt.Println(resultado)

	multiplica := func(x int) int {
		return x * 2
	}

	resultado2 := multiplica(8)

	fmt.Println(resultado2)
}

func soma(a, b int) int {
	return a + b
}
