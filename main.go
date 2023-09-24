package main

import "fmt"

func main() {
	var m map[string]int
	fmt.Println(m)

	m = map[string]int{"foo": 1, "bar": 2}
	fmt.Println(m)

	fmt.Println(m["foo"])
	m["bar"] = 88

	delete(m, "foo")
	m["baz"] = 784
	fmt.Println(m)

	fmt.Println(m["foo"])
	v, ok := m["foo"]
	fmt.Println(v, ok);
}
