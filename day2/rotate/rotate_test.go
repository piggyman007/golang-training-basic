package rotate

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name           string
		input          []int
		expected []int
	}{
		{
			name:			"Test#1",
			input:		[]int{1,2,3,4,5},
			expected: []int{5,1,2,3,4},
		},
		{
			name: 		"Test#2",
			input:    []int{1,2,3},
			expected: []int{3,1,2},
		},
		{
			name: 		"Test#3",
			input:    []int{},
			expected: []int{},
		},
		{
			name: 		"Test#4",
			input:    []int{5},
			expected: []int{5},
		},
		{
			name: 		"Test#5",
			input:    []int{1,2},
			expected: []int{2,1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rotate(tt.input)
			if !reflect.DeepEqual(r, tt.expected) {
				t.Errorf("expect % #v but got % #v", tt.expected, r)
			}
		})
	}
}