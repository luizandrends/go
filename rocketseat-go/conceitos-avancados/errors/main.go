package main

import (
	"errors"
	"fmt"
)

func main() {
	user, err := NewUser(true)

	if err != nil {
		fmt.Println("Algum error na hora de criar o usuário")
		return
	}

	user.Foo()
}
type User struct {
	foo string
}

func (u User) Foo() {
	fmt.Println(u.foo)
}

func NewUser(wantErr bool) (*User, error) {
	if wantErr {
		return nil, errors.New("Um Erro")
	}
	return &User{}, nil
}
