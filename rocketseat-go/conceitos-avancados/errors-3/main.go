package main

import (
	"errors"
	"fmt"
)

func main() {
	err := foo()
	fmt.Println(err)
	fmt.Println(errors.Is(err, ErrQualquer))
	fmt.Println(errors.Is(err, ErrQualquer2))
}

var ErrQualquer = errors.New("Error")
var ErrQualquer2 = errors.New("Error 2")

func a() error {return ErrQualquer}
func b() error {return ErrQualquer2}

func foo() error {return ErrQualquer}