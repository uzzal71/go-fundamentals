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

type menu []menuItem

func (m menu) print() {
	for _, item := range m {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", len(item.name)))
		for size, price := range item.prices {
			fmt.Printf("\t%8s%10.2f\n", size, price)
		}
	}
}

func (m *menu) add() {
	fmt.Println("Please enter the name of the new item:")
	name, _ := in.ReadString('\n')
	name = strings.TrimSpace(name)
	data = append(*m, menuItem{name: name, prices: map[string]float64{}})
}

var in = bufio.NewReader(os.Stdin)

func AddItem() {
	data.add()
}

func PrintMenu() {
	data.print()
}