package unsplice

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
			name:           "Test#1",
			input:          "ab\\\ncd\\\nef",
			expectedResult: "abcdef",
		},
		{
			name:           "Test#2",
			input:          "abc\\\ndef",
			expectedResult: "abcdef",
		},
		{
			name:           "Test#3",
			input:          "abc\ndef",
			expectedResult: "abc\ndef",
		},
		{
			name:           "Test#4",
			input:          "abc\\def",
			expectedResult: "abc\\def",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := run(tt.input)
			if result != tt.expectedResult {
				t.Error("case: "+tt.name+", expect: ", tt.expectedResult)
			}
		})
	}
}
