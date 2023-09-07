package main

import (
	"fmt"
)

func main() {
	myName  := "Uzzal Roy"
	amount := 2
	totalAmount := float32(amount)

	fmt.Println(myName)
	fmt.Println(totalAmount)

	a, b := 10, 8
	c := a + b

	fmt.Println(c)

	e := a == b 
	
	fmt.Println(e)
	fmt.Println(a-b)
	fmt.Println(a*b)
	fmt.Println(a/b)
	fmt.Println(a%b)

	fmt.Println(float32(a) / float32(b))
}