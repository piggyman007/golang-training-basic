package unsplice

import (
	"fmt"
	"strings"
)

func run(input string) string {
	lines := strings.Split(input, "\n")
	res := ""

	if len(lines) == 1 {
		res = lines[0]
	} else {
		for _, line := range lines {
			idx := strings.LastIndex(line, "\\")
			if idx != -1 {
				res += line[0:idx] + line[idx+1:]
			} else {
				res += line + "\n"
			}
		}
	}

	respDisp := strings.TrimSpace(res)
	fmt.Println(respDisp)

	return respDisp
}
