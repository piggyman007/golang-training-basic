package lcddigit

import (
	"fmt"
	"strconv"
	"strings"
)

var Nums = []string{
	"._.\n|.|\n|_|",
	"...\n..|\n..|",
	"._.\n._|\n|_.",
	"._.\n._|\n._|",
	"...\n|_|\n..|",
	"._.\n|_.\n._|",
	"._.\n|_.\n|_|",
	"._.\n..|\n..|",
	"._.\n|_|\n|_|",
	"._.\n|_|\n..|",
}

func run(n string) string {
	res := make([]string, 0)
	for _, c := range n {
		i, err := strconv.Atoi(string(c))
		if err != nil {
			panic(err)
		}
		res = append(res, Nums[i])
	}

	resStr := strings.Join(res, "")
	print(res)

	return resStr
}

func print(res []string) {
	resLine1 := make([]string, 0)
	resLine2 := make([]string, 0)
	resLine3 := make([]string, 0)

	for _, r := range res {
		lines := strings.Split(r, "\n")
		resLine1 = append(resLine1, lines[0])
		resLine2 = append(resLine2, lines[1])
		resLine3 = append(resLine3, lines[2])
	}

	fmt.Println(strings.Join(resLine1, "\t"))
	fmt.Println(strings.Join(resLine2, "\t"))
	fmt.Println(strings.Join(resLine3, "\t"))
}
