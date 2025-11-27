package analyzer

import (
	"fmt"
)

type Stats struct {
	Total int
	Info  int
	Warn  int
	Error int
}

var total, info, warn, error int

func CalculateStats(filename string) {

	lines := ParseFile(filename)
	total = len(lines)
	fmt.Println("Количество записей: ", total)
	fmt.Println("Все записи: ", lines)
}
