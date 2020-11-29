package main

import (
	"errors"
	"fmt"
)

func suma(a, b int) (int, error) {
	if a < b {
		return 0, errors.New("el primer valor es menor que el segundo a < b")
	}
	return a + b, nil
}

func main() {
	fmt.Println("Hello world")
	r, err := suma(1, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}
