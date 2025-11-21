package analyzer

import (
	"fmt"
	"strings"
)

func Stats(param string) {
	line := Parser("")
	fmt.Printf(strings.Join(line, param))
}
