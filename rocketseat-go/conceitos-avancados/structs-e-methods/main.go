package main

import (
	"encoding/json"
	"fmt"
)

type MinhaString string

type User struct {
	// foo.Foo
	Name string `json:"name"`
	ID   uint64 `json:"id"`
}

func (u User) PrintName() {
	fmt.Println(u.Name)
}

func (u *User) UpdateName(newName string) {
	u.Name = newName
}

func main() {
	user := &User{"Pedro Pessoa", 10}
	// anotherUser := User{Name: "Pedro Pessoa", ID: 10}
	res, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	// user.Bar()
}