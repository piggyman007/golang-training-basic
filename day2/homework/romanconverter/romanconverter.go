package romanconverter

var decimalValue = []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
var romanNumeral = []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

// Convert integer to roman number
func Convert(n int) string {
	r := ""
	for index := 0; index < len(decimalValue); index++ {
		for decimalValue[index] <= n {
			r += romanNumeral[index]
			n -= decimalValue[index]
		}
	}

	return r
}
