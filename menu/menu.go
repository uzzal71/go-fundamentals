package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type menuItem struct {
	name   string
	prices map[string]float64
}

var in = bufio.NewReader(os.Stdin)

func AddItem() {
	fmt.Println("Please enter the name of the new item:")
	name, _ := in.ReadString('\n')
	name = strings.TrimSpace(name)
	data = append(data, menuItem{name: name, prices: map[string]float64{}})
}

func PrintMenu() {
	for _, item := range data {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", len(item.name)))
		for size, price := range item.prices {
			fmt.Printf("\t%8s%10.2f\n", size, price)
		}
	}
}