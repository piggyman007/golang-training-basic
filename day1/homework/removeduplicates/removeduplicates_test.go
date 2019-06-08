package removeduplicates

import (
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name           string
		input          []int
		expectedResult []int
	}{
		{
			name:           "Test#1",
			input:          []int{3, 1, 2, 2},
			expectedResult: []int{1, 2, 3},
		},
		{
			name:           "Test#2",
			input:          []int{1,2},
			expectedResult: []int{1,2},
		},
		{
			name:           "Test#3",
			input:          []int{1,1,2,2,3,3},
			expectedResult: []int{1,2,3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := run(tt.input)
			if !testEq(result, tt.expectedResult) {

				t.Error("case: "+tt.name+", expect: ", tt.expectedResult)
			}
		})
	}
}

func testEq(a, b []int) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
