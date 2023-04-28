package main

import (
	"errors"
	"fmt"
)

func main() {
	s1 := "GGACGGATTCTG"
	s2 := "AGGACGGATTCT"
	di, _ := Distance(s1, s2)
	fmt.Println(di)
}

func Distance(a, b string) (int, error) {
	count := 0
	if a == "" && b == "" {
		return 0, errors.New("empty string")
	}
	if len(a) > len(b) {
		return 0, nil
	}
	if len(a) < len(b) {
		return 0, nil
	}
	if len(a) == 0 {
		return 0, errors.New("empty string")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}
