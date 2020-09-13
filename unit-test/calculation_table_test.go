package main

import "testing"

func TestSums(t *testing.T) {
	tables := []struct {
		x float64
		y float64
		n float64
	}{
		{1.0, 1.0, 2.0},
		{1.0, 2.0, 3.0},
		{2.0, 2.0, 4.0},
		{5.0, 2.0, 7.0},
	}
	t.Log("Given the need to test sum operation.")
	{
		for _, table := range tables {
			t.Logf("When sum of %g and %g", x, y)
			{
				total := Sum(table.x, table.y)
				if total != table.n {
					t.Errorf("Should sum of (%g+%g) is incorrect, got: %g, want: %g.", table.x, table.y, total, table.n)
				}
			}
		}
	}
}