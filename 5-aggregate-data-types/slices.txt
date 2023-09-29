package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

func main() {
	var s []int 
	fmt.Println(s)

	s = []int{1, 2, 3}
	fmt.Println(s)

	s[1] = 99
	fmt.Println(s)

	s = append(s, 5, 14, 25)
	fmt.Println(s)

	s = slices.Delete(s, 1, 3)
	fmt.Println(s)
}
