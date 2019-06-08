package calcstats

import (
	"math"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name           string
		input          []int
		expectedResult Stat
	}{
		{
			name:           "Test#1",
			input:          []int{6, 9, 15, -2, 92, 11},
			expectedResult: Stat{min: -2, max: 92, no: 6, avg: 21.833333},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := run(tt.input)
			if result.min != tt.expectedResult.min ||
				result.max != tt.expectedResult.max ||
				result.no != tt.expectedResult.no ||
				(math.Round(result.avg*10000)/10000) != (math.Round(tt.expectedResult.avg*10000)/10000) {
				t.Error("case: "+tt.name+", expect: ", tt.expectedResult)
			}
		})
	}
}
