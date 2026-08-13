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
	percentaje float64
	bar        string
	who        string
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

	who := "Maxi"

	for scanner.Scan() {
		line := scanner.Text()

		if igText, match := matchIgnore(line); match {
			if igText == "Consumos Tarjeta:4824690004612580" {
				who = "Cele"
			}
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

		key := matchKey(name, who)
		cat := matchCategory(strings.Replace(key, " - "+who, "", 1))

		_, ok := moneyMap[key]
		if !ok {
			moneyMap[key] = &Item{
				name:     key,
				amount:   amount,
				category: cat,
				who:      who,
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
		row.percentaje = 100.0 / float64(total) * float64(row.amount)
		row.bar = strings.Repeat("=", max(int(row.percentaje), 0)) + ">"
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

func padRight(str string, target int) string {
	spaces := target - len(str)

	return fmt.Sprintf("%s%s", str, strings.Repeat(" ", spaces))
}
