package main

import (
	"errors"
	"fmt"
	"math"
)

type SqrtError struct {
	msg string
}

func (s SqrtError) Error() string { return s.msg }

func raizQuadrada(x float64) (float64, error){
	if x < 0 {
		return 0, SqrtError{"não existe raiz quadrada de numero negativo"}
	}
	resultado := math.Sqrt(x)
	return resultado, nil
}

var ErrorNotFound = errors.New("Not found")

func main() {
	// err := foo()
	// if err != nil && errors.Is(err, ErrorNotFound) {
	// 	fmt.Println("Deu o erro not found")
	// 	return
	// }

	// fmt.Println("Foi pra fora")

	err := foo()
	var sqrtError *SqrtError
	if err != nil && errors.As(err, &sqrtError) {
		fmt.Println(sqrtError.msg)
		return
	}
}

func foo() error { return SqrtError{msg: "teste"} }
