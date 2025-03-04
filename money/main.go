package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	// moneyMap, catMap, total := loadFile("mov visa.csv")
	moneyMap, catMap, total, padMax := loadFile("visa.csv")

	var items []*Item
	for item := range maps.Values(moneyMap) {
		items = append(items, item)
	}

	slices.SortFunc(items, func(a *Item, b *Item) int {
		return b.amount - a.amount
	})

	for _, row := range items {
		fmt.Printf("%s : %d  --  %d%% ||| %s\n", padName(row.name, padMax), row.amount, row.percentaje, row.bar)
	}
	fmt.Println()
	fmt.Println("TOTAL: ", total)

	fmt.Println()
	fmt.Println("BY CATEGORY")

	catSum := 0
	for key, value := range catMap {
		fmt.Printf("%s: %d\n", key, value)
		catSum += value
	}
	fmt.Println("----------")
	fmt.Printf("%s: %d\n", "TOTAL CATS", catSum)
	fmt.Printf("%s: %d\n", "RESTO", total-catSum)

}
