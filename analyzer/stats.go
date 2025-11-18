package analyzer

import (
	"fmt"
	"strings"
)

func stats(lines []string) {
	fmt.Printf(strings.Join(lines, "[INFO]"))
}
