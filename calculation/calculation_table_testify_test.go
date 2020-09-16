package calculation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumsWithTestify(t *testing.T) {
	tables := []struct {
		x float64
		y float64
		expected float64
	}{
		{1.0, 1.0, 2.0},
		{1.0, 2.0, 3.0},
		{2.0, 2.0, 4.0},
		{5.0, 2.0, 7.0},
	}
	for _, table := range tables {
		assert.Equal(t, Sum(table.x, table.y), table.expected)
	}
}