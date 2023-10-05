package main

import (
	"fmt"
	"strings"
)

type printer interface {
	Print() string
}

type user struct {
	username string 
	id int
}

func(u user) Print() string {
	return fmt.Printf("%v [%v]\n", u.username, u.id)
}

type menuItem(mi menuItem) Print() string {
	var b bytes.Buffer 
	b.WriteString(mi.name + "\n")
	b.WriteString(strings.Repeat("-", 10) + "\n")
	for size, cost := range mi.prices {
		fmt.Fprintf(&b, "\t%10s10.2f\n", size, cost)
	}
	return b.String()
}

func main() {

}