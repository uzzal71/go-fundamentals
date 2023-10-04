package main

func main() {
	
}

type user struct {
	id 			int
	username	string
}

func (u user) String() string {							// value receiver
	return fmt.Printf("%v (%v)\n", u.username, u.id)
}

func (u *user) UpdateName(n name) {
	u.username = name
}