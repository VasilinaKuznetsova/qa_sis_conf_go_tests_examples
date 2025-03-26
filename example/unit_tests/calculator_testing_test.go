package example

import "testing"

func TestSum(t *testing.T) {
	a, b := 2, 3
	want := 5
	result := Sum(a, b)

	if result != want {
		t.Errorf("Summ(%d, %d) = %d; ожидается: %d", a, b, result, want)
	}
}
