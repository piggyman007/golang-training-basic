package phoneNumberNormalizer

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNormalize(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected map[string]int
	}{
		{
			name: "Test#1",
			input: []string{
				"1234567890",
				"123 456 7891",
				"(123) 456 7892",
				"(123) 456-7893",
				"123-456-7894",
				"123-456-7890",
				"1234567892",
				"(123)456-7892",
			},
			expected: map[string]int{
				"1234567890": 2,
				"1234567891": 1,
				"1234567892": 3,
				"1234567893": 1,
				"1234567894": 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Normalize(tt.input)
			fmt.Println(">>>> ", r)
			if !reflect.DeepEqual(r, tt.expected) {
				t.Errorf("expect % #v but got % #v", tt.expected, r)
			}
		})
	}
}
