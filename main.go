package main

import "fmt"

func main() {
	// for { ... } infinite loop
	// for condition { ... } loop till condition
	// for initializer; test; post clause { ... } counter-base loop
	// Infinite Loops
	/*
	i := 1
	for {
		fmt.Println(i)
		i += 1
	}
	*/

	// Loop till condition
	/*
	i := 1
	for i < 3 {
		fmt.Println(i)
		i += 1
	}
	fmt.Println("Done!")
	*/

	// Counter-based Loops
	for i:=1; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Println("Done!")
}
