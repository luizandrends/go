package main

import "fmt"

func main() {
	ShareWith("")
}

func ShareWith(name string) string {
	dialogue := ""

	if name != "" {
		dialogue = "One for " + name + ", one for me."
	} else {
		dialogue = "One for you, one for me."
	}

	fmt.Println(dialogue)
	return dialogue
}