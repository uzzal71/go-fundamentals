package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type menuItem struct {
	name string
	prices map[string]float64
}

menu := []menuItem{
	{name: "Coffee", prices: map[string]float64{"small": 1.54, "medium": 1.74, "large": 1.98}},
	{name: "Espresso", prices: map[string]float64{"small": 1.25, "medium": 2.74, "tiple": 1.51}},
}

in := bufio.NewReader(os.Stdin)

func main() {
	
	loop:
	for {
		fmt.Println("Please select a option")
		fmt.Println("1) Print menu")
		fmt.Println("2) Add item")
		fmt.Println("q) quit")
		choice, _ := in.ReadString('\n')
		
		switch strings.TrimSpace(choice) {
		case "1":
			for _, item := range menu {
				fmt.Println(item.name)
				fmt.Println(strings.Repeat("-", 10))
				for size, price := range item.prices {
					fmt.Printf("\t%8s%10.2f\n", size, price)
				}
			}
		case "2":
			fmt.Println("Please enter the name of the new item")
			name, _ := in.ReadString('\n')
			menu = append(menu, menuItem{name:name, prices: map[string]float64{}})
		case "q":
			break loop
		default:
			fmt.Println("Unknown option")
		}
	}
}
