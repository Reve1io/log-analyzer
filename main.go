package main

import (
	"log-analyzer/analyzer"
	"os"
)

func main() {
	filename := os.Args[1]
	analyzer.Parser(filename)
}
