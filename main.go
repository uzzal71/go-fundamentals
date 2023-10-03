package main

func main() {
	
}

// function
var i int 

func isEvent(i int) bool {
	return  i% 2 == 0
}
ans := isEven(i)

// method
type myInt int
var mi myInt
func (i myInt) isEven() bool {
	return int(i) % 2 == 0
}
ans = mi.isEven()