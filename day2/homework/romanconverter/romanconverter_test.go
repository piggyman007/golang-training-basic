package romanconverter

import (
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{name: "Test#1", input: 1, expected: "I"},
		{name: "Test#2", input: 4, expected: "IV"},
		{name: "Test#3", input: 5, expected: "V"},
		{name: "Test#4", input: 6, expected: "VI"},
		{name: "Test#5", input: 27, expected: "XXVII"},
		{name: "Test#6", input: 58, expected: "LVIII"},
		{name: "Test#7", input: 100, expected: "C"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Convert(tt.input)
			if !reflect.DeepEqual(r, tt.expected) {
				t.Errorf("expect % #v but got % #v", tt.expected, r)
			}
		})
	}
}
