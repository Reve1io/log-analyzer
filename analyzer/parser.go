package analyzer

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type LogEntry struct {
	Timestamp string
	Level     string
	Message   string
}

func ParseFile(filename string) []LogEntry {

	path := filepath.Join("logs", filename)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Ошибка! Файл отсутсвует", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	maxCapacity := 1 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	var logs []LogEntry

	for scanner.Scan() {
		line := scanner.Text()

		start := strings.Index(line, "[")
		end := strings.Index(line, "]")

		var ts, lvl, msg string

		if start != -1 && end != -1 && end > start {
			ts = strings.TrimSpace(line[:start])
			lvl = strings.TrimSpace(line[start+1 : end])
			msg = strings.TrimSpace(line[end+1:])
		} else {
			msg = line
		}

		logs = append(logs, LogEntry{
			Timestamp: ts,
			Level:     lvl,
			Message:   msg,
		})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка сканирования файла:", err)
		return logs
	}

	return logs
}
