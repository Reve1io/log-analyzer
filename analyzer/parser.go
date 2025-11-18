package analyzer

import (
	"bufio"
	"fmt"
	"os"
)

func Parser(filename string) []string {
	filepath := "logs\\" + filename

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Ошибка! Файл отсутсвует", err)
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	maxCapacity := 1 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	for scanner.Scan() {
		lines := append(lines, scanner.Text())
		fmt.Println(lines)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка сканирования файла:", err)
		return lines
	}

	return lines
}
