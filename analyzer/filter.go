package analyzer

import (
	"fmt"
)

func FilteredFile(filename, flag string) []string {
	filteredLines := []string{}
	lines := ParseFile(filename)

	for _, line := range lines {
		if line.Level == flag {
			filteredLines = append(filteredLines, line.Timestamp, line.Level, line.Message)
		}
	}

	if len(filteredLines) == 0 {
		fmt.Printf("Записи с параметром %s отсутствуют, попробуйте debug, info или error\n", flag)
	}

	return filteredLines
}
