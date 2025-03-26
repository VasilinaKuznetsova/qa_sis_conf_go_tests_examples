package example

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumTestify(t *testing.T) {
	a, b := 2, 3
	want := 5
	result := Sum(a, b)
	assert.Equal(t, want, result, fmt.Sprintf("Summ(%d, %d) = %d; ожидается: %d", a, b, result, want))
}
