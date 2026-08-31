package main

import (
	"errors"
)

func main() {
	Distance("AATG", "AAA")
}

func Distance(a, b string) (int, error) {
	var diff int

	if len(a) != len(b) {
		return 0, errors.New("Distance error")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff += 1
		}
	}

	return diff, nil
}
