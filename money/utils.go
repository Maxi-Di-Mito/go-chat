package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	name       string
	amount     int
	category   string
	percentaje int
	bar        string
}

func loadFile(fileName string) (map[string]*Item, map[string]int, int, int) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	moneyMap := make(map[string]*Item)
	catMap := make(map[string]int)

	total := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if matchIgnore(line) {
			continue
		}

		parts := strings.Split(line, ";")
		// date := parts[0]
		name := parts[1]
		// id := parts[2]
		amount := parseAmount(parts[3])
		// amountUSS := parseAmount(parts[4])
		if amount == 0 {
			continue
		}

		key := matchKey(name)
		cat := matchCategory(key)

		_, ok := moneyMap[key]
		if !ok {
			moneyMap[key] = &Item{
				name:     key,
				amount:   amount,
				category: cat,
			}
		} else {
			moneyMap[key].amount += amount
		}

		if cat != "" {
			_, okCat := catMap[cat]
			if !okCat {
				catMap[cat] = amount
			} else {
				catMap[cat] += amount
			}
		}

		total += amount

	}

	var maxname int

	for _, row := range moneyMap {
		row.percentaje = int(100.0 / float64(total) * float64(row.amount))
		row.bar = strings.Repeat("=", row.percentaje) + ">"
		if len(row.name) > maxname {
			maxname = len(row.name)
		}
	}

	return moneyMap, catMap, total, maxname
}

func parseAmount(str string) int {
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(err)
	}
	return int(value)
}

func padName(name string, target int) string {
	spaces := target - len(name)

	return fmt.Sprintf("%s%s", name, strings.Repeat(" ", spaces))
}
