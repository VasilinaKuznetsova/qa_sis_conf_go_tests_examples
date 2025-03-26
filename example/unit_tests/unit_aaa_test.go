package example

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumAAA(t *testing.T) {
	t.Parallel()

	t.Run("positive case: sum of positive numbers", func(t *testing.T) {
		//Arrange
		a, b := 2, 3
		want := 5
		t.Log(fmt.Sprintf("Input values: %d, %d; expected value: %d", a, b, want))

		//Act
		result := Sum(a, b)

		//Assert
		assert.Equal(t, want, result, fmt.Sprintf("Summ(%d, %d) = %d; expected: %d", a, b, result, want))
	})
}
