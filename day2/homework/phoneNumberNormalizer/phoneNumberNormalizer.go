package phoneNumberNormalizer

import (
	"strings"
)

// Normalize phone number
func Normalize(nos []string) map[string]int {
	r := map[string]int{}
	for _, no := range nos {
		decoratedNo := strings.Replace(no, " ", "", -1)
		decoratedNo = strings.Replace(decoratedNo, "(", "", -1)
		decoratedNo = strings.Replace(decoratedNo, ")", "", -1)
		decoratedNo = strings.Replace(decoratedNo, "-", "", -1)

		if _, ok := r[decoratedNo]; ok {
			r[decoratedNo]++
		} else {
			r[decoratedNo] = 1
		}
	}

	return r
}
