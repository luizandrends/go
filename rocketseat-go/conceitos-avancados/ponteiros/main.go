package main

import "fmt"

func main() {
	// x := 10
	// take(&x)
	// fmt.Println(x)

	x := create()
	fmt.Println(*x)
}

func take(x *int) {
	*x = 100
}

func create() *int {
	x := 10
	return &x
}

func foo(x *int) {
	*x = 100
}