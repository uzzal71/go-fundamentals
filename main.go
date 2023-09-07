package main

import (
	"fmt"
)

func main() {
	/*
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
	
	const a = 45
	var b float32 = 45
	var f32 float32 = b 
	var f64 float64 = float64(f32)

	fmt.Println(f32, f64)

	const c = iota 
	fmt.Println(c)

	const (
		d = 2 * 5
		e // = 2 * 5
		f = iota
		g
		h = 10 * iota
	)

	fmt.Println(d, e, f, g, h)
	*/

	s := "Hello, world!"

	p := &s
	fmt.Println(p)
	fmt.Println(*p)

	*p = "hello ,gopheres!"
	fmt.Println(s)

	p = new(string)

	fmt.Println(p, *p)
}