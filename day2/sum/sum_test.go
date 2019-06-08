package sum

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	fmt.Println(x)
	fmt.Println(Y)
	if sum(5, 5) != 10 {
		t.Errorf("Error %v\n", 10)
	}

	if sum(4, 3) != 7 {
		t.Errorf("Error %v\n", 7)
	}
}