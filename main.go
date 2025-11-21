package main

import (
	"log-analyzer/analyzer"
	"os"
)

func main() {
	filename := os.Args[1]
	param := os.Args[2]
	analyzer.Parser(filename)
	analyzer.Stats(param)
}
