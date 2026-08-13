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

	var maxAmount int
	for _, row := range items {
		if l := len(ac.FormatMoney(row.amount)); l > maxAmount {
			maxAmount = l
		}
	}

	for _, row := range items {
		fmt.Printf("%s : %s  --  %.1f%% ||| %s  : %s\n", padRight(row.name, padMax), padRight(ac.FormatMoney(row.amount), maxAmount), row.percentaje, row.bar, row.who)
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
