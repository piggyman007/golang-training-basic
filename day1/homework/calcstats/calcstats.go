package calcstats

type Stat struct {
	min int
	max int
	no  int
	avg float64
}

func run(input []int) Stat {
	min := 0
	max := 0
	no := len(input)
	sum := float64(0.0)
	for _, num := range input {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
		sum += float64(num)
	}

	res := Stat{
		min: min,
		max: max,
		no:  no,
		avg: sum / float64(no),
	}

	return res
}
