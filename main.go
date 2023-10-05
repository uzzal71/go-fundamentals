package main

type printer interface {
	Print() string
}

type user struct {
	username string 
	id int
}

func(u user) Print() string {
	
}

func main() {

}