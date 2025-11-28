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

func CalculateStats(filename, flag string) {

	lines := ParseFile(filename)
	filter := FilteredFile(filename, flag)

	total := len(lines)

	sliceInfo, sliceWarn, sliceError := []string{}, []string{}, []string{}

	for _, line := range lines {
		if line.Level == "INFO" {
			sliceInfo = append(sliceInfo, line.Timestamp, line.Level, line.Message)
		} else if line.Level == "DEBUG" {
			sliceWarn = append(sliceWarn, line.Timestamp, line.Level, line.Message)
		} else if line.Level == "ERROR" {
			sliceError = append(sliceError, line.Timestamp, line.Level, line.Message)
		}
	}

	fmt.Println("Количество записей: ", total)
	fmt.Println("Количество информационных записей: ", len(sliceInfo))
	fmt.Println("Количество рабочих записей: ", len(sliceWarn))
	fmt.Println("Количество записей с ошибками: ", len(sliceError))

	if flag != "" {
		fmt.Printf("Все записи c параметром %v: %v\n", flag, filter)
	}
}
