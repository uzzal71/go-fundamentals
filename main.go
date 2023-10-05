package main

import (
	"fmt"
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

type menuItem(m menuItem) Print() string {
	
}

func main() {

}