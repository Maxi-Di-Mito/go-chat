package main

import (
	"regexp"
	"strings"
)

var ignores []string = []string{"Fecha", "Consumos", "Tarjeta:", "BONIF", "SU PAGO"}

func matchIgnore(line string) bool {
	for _, ig := range ignores {
		if strings.Contains(line, ig) {
			return true
		}
	}
	return false
}

var matchersKeys map[string]string = map[string]string{
	"SWISS":            "Swiss",
	"PEDIDOS":          "PEDIDOSYA",
	"Microsoft":        "XBOX",
	"MICROSOFT":        "XBOX",
	"COTO":             "COTO",
	"FEDERACION":       "SEGURO AUTO",
	"ZURICH":           "SEGURO/RETIRO",
	"MCDONALDS":        "MCDonalds",
	"KFC":              "KFC",
	"MOSTAZA":          "Mostaza",
	"BURGERKING":       "BURGER",
	"INST PEDRO GIACH": "COLEGIO",
	"ESUR":             "LUZ",
	"MTGA":             "GAS",
	"CHANGO":           "Changomas",
	"HOYTS":            "HOYTS",
	"SHOWCASE":         "SHOWCASE",
	"UBER":             "UBER",
	"YPF":              "YPF",
	"LIAN":             "CHINO",
	"CABIFY":           "UBER",
	"SUPERDIA":         "DIA",
	"PERSONAL":         "INTERNET/CEL",
	"FARMA":            "FARMACIA",
}

var categories map[string]string = map[string]string{
	"MCDonalds":    "FAST FOOD",
	"KFC":          "FAST FOOD",
	"Mostaza":      "FAST FOOD",
	"BURGER":       "FAST FOOD",
	"LUZ":          "Servicios",
	"GAS":          "Servicios",
	"INTERNET/CEL": "Servicios",
	"COTO":         "Supermercado",
	"Changomas":    "Supermercado",
	"CHINO":        "Supermercado",
	"DIA":          "Supermercado",
	"HOYTS":        "CINE",
	"SHOWCASE":     "CINE",
	"PEDIDOSYA":    "DELIVERY",
}

func matchKey(line string) string {
	for key, value := range matchersKeys {
		if strings.Contains(line, key) {
			return value
		}
		if match, _ := regexp.MatchString("\\d\\d/\\d\\d", line); match {
			return "CUOTAS"
		}
	}
	return line
}

func matchCategory(name string) string {
	cat, ok := categories[name]
	if ok {
		return cat
	} else {
		return ""
	}

}
