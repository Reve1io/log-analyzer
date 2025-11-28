package main

import (
	"fmt"
	"log-analyzer/analyzer"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Укажите название лог файла, например:")
		fmt.Println("log-analyzer app.log info")
		return
	}

	filename := os.Args[1]

	var flag string
	if len(os.Args) >= 3 {
		flag = strings.ToUpper(os.Args[2])
	}

	analyzer.CalculateStats(filename, flag)
}
