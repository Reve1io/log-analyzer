package main

import (
	"log-analyzer/analyzer"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]
	flag := strings.ToUpper(os.Args[2])

	analyzer.CalculateStats(filename, flag)
}
