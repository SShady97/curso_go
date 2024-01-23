package main

import (
	"errors"
	"fmt"
)

func Distance(a, b string) (int, error) {
	count := 0
	if len(a) != len(b) {
		return 0, errors.New("Ha ocurrido un error")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count += 1
		}
	}
	return count, nil
}

func main() {
	s1 := "ABCDEG"
	s2 := "ABCDFG"
	count, _ := Distance(s1, s2)
	fmt.Println(count)
}
