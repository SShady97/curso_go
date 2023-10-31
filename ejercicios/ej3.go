package main

import (
	"errors"
	"fmt"
)

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("No puedes ingresar el valor menor a 1")
	}
	if n == 1 {
		return 0, nil
	}

	var count int

	for count = 0; n > 1; count++ {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
	}
	return count, nil
}

func main() {
	fmt.Println(CollatzConjecture(12))
}
