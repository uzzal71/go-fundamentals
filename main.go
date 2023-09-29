package main

import "fmt"

func main() {
	// for { ... } infinite loop
	// for condition { ... } loop till condition
	// for initializer; test; post clause { ... } counter-base loop
	// Infinite Loops
	i := 1
	for {
		fmt.Println(i)
		i += 1
	}
}
