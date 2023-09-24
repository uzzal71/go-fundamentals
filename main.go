package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Please select a option")
	fmt.Println("1) Print menu")
	in := bufio.NewReader(os.Stdin)
	choice, _ := in.ReadString('\n')
	choice = strings.TrimSpace(choice) // we don't know what to do with this yet!

	type menuItem struct {
		name string
		prices map[string]float64
	}

	menu := []menuItem{
		{name: "Coffee", prices: map[string]float64{"small": 1.54, "medium": 1.74, "large": 1.98}},
		{name: "Espresso", prices: map[string]float64{"small": 1.25, "medium": 2.74, "tiple": 1.51}},
		{name: "Chappission", prices: map[string]float64{"small": 2.54, "medium": 3.74, "large": 1.98}},
	}

	fmt.Println(menu)
}
