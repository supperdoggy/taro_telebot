package main

import (
	"fmt"
)

var (
	mast = map[string]string{
		"p": "пентаклей",
		"s": "мечей",
		"w": "жезлов",
		"c": "кубков",
	}
	val = map[string]string{
		"ace":   "Туз",
		"king":  "Король",
		"queen": "Королева",
		"page":  "Паж",
		"2":     "Двойка",
		"3":     "Тройка",
		"4":     "Четверка",
		"5":     "Пятерка",
		"6":     "Шестерка",
		"7":     "Семерка",
		"8":     "Восьмерка",
		"9":     "Девятка",
		"10":    "Десятка",
	}
)

func main() {
	for k, v := range val {
		for i, j := range mast {
			fmt.Printf("\"%s_%s\": \"%s %s\",\n", k, i, v, j)
		}

	}
}