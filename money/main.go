package main

import (
	"fmt"
	"maps"
	"slices"

	"github.com/leekchan/accounting"
)

func main() {
	// moneyMap, catMap, total := loadFile("mov visa.csv")
	moneyMap, catMap, total, padMax := loadFile("movimientos.csv")

	var items []*Item
	for item := range maps.Values(moneyMap) {
		items = append(items, item)
	}

	slices.SortFunc(items, func(a *Item, b *Item) int {
		return b.amount - a.amount
	})

	ac := accounting.Accounting{Symbol: "$", Precision: 0, Thousand: ".", Decimal: ","}

	for _, row := range items {
		fmt.Printf("%s : %s  --  %d%% ||| %s\n", padName(row.name, padMax), ac.FormatMoney(row.amount), row.percentaje, row.bar)
	}
	fmt.Println()
	fmt.Println("TOTAL: ", ac.FormatMoney(total))

	fmt.Println()
	fmt.Println("BY CATEGORY")

	catSum := 0
	for key, value := range catMap {
		fmt.Printf("%s: %s\n", key, ac.FormatMoney(value))
		catSum += value
	}
	fmt.Println("----------")
	fmt.Printf("%s: %s\n", "TOTAL CATS", ac.FormatMoney(catSum))
	fmt.Printf("%s: %s\n", "RESTO", ac.FormatMoney(total-catSum))

}
