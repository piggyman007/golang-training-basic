package removeduplicates

import (
	"sort"
)

func run(nums []int) []int {
	set := map[int]bool{}
	res := make([]int, 0)

	for _, n := range nums {
		if _, ok := set[n]; !ok {
			set[n] = true
		}
	}

	for n := range set {
		res = append(res, n)
	}

	sort.Ints(res)

	return res
}
