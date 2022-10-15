package operations

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	suite := []struct {
		num1, num2 int
		expected   int
	}{
		{1, 2, 3},
		{2, -2, 0},
		{11, 2, 13},
		{0, 0, 0},
	}

	for _, v := range suite {
		if sum := add(v.num1, v.num2); sum != v.expected {
			t.Errorf("Failed! Wants %d, but got %d", v.expected, sum)
		} else {
			t.Logf("Succeeded! Wants %d, and got %d", v.expected, sum)

		}
	}
}

func ExampleAdd() {
	fmt.Println(add(2, 3))
	// Output:
	// 5
}

func BenchmarkAdd(b *testing.B) {

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		add(2, 3)
	}
}
