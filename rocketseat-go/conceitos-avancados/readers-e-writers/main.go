package main

import (
	"errors"
	"fmt"
)
var ErrNotFound = errors.New("not found")
func main() {
 err := fmt.Errorf("wrap: %w", ErrNotFound)
 if errors.Is(err, ErrNotFound) {
  fmt.Println("Error is ErrNotFound")
 }
}