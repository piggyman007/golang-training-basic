package rotate

// Rotate slice of []int
func Rotate(n []int) []int {
	size := len(n)
	r := make([]int, size)
	if size == 0 || size == 1 {
		return n
	}
	r[0] = n[size - 1]
	for i := 0; i < size - 1; i++ {
		r[i+1] = n[i]
	}

	return r
}

