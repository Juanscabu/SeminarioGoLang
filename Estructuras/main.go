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
	fmt.Println("--------Arreglos----------")
	var arr [2]int
	arr[0] = 1
	arr[1] = 2
	for i, v := range arr {
		fmt.Println(i, v)
	}

	fmt.Println("-------Slices-----------")

	//var l []int
	l := make([]int, 10)
	l = append(l, 10)
	l = append(l, 20)
	l = append(l, 30)

	for i, v := range l {
		fmt.Println(i, v)
	}

	fmt.Println("-------Maps-----------")
	m := make(map[int]string)
	m[0] = "a"
	m[1] = "b"

	for k, v := range m {
		fmt.Println(k, v)
	}

}
