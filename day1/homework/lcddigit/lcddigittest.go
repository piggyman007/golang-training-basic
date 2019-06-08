package lcddigit

import (
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedResult string
	}{
		{
			name:           "Test 191",
			input:          "191",
			expectedResult: Nums[1] + Nums[9] + Nums[1],
		},
		{
			name:           "Test 910",
			input:          "910",
			expectedResult: Nums[9] + Nums[1] + Nums[0],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := run(tt.input)
			if result != tt.expectedResult {
				t.Error("case: " + tt.name + ", expect: " + tt.expectedResult)
			}
		})
	}
}
