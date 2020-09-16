package calculation

import (
	"testing"
)

var (
	x = 25.0
	y = 4.0
	expectedSum = 29.0
	expectedDivide = 6.25
)
//  Given, When, Should format

func TestSum(t *testing.T){
	t.Log("Given the need to test sum operation.")
	{
		t.Logf("When sum of %g and %g", x, y)
		{
			total := Sum(x, y)
			if total != expectedSum {
				t.Errorf("Should sum result is incorrect, got: %g, want: %g.", total, expectedSum)
			} else {
				t.Logf("Should sum result is correct, got: %g , want: %g.", total, expectedSum)
			}
		}
	}
}

func TestDivide(t *testing.T){
	t.Log("Given the need to test divide operation.")
	{
		t.Logf("When  divident %g with divider is non zero %g " ,x,  y)
		{
			divide, _ := Divide(x, y)

			if divide != expectedDivide {
				t.Errorf("Should division result is incorrect, got: %g, want: %g.", divide, expectedDivide)
			} else {
				t.Logf("Should division result is correct, got: %g , want: %g.", divide, expectedDivide)
			}
		}

		t.Logf("When divident %g with divider is zero %g " ,x, y)
		{
			y = 0
			divide,err := Divide(x, y)

			if err == nil {
				t.Errorf("Should error(%v) is not nil , equal: %v.", err,  DivideByZeroError )
			} else{
				t.Logf("Should error(%v) is correct value %v .", err, DivideByZeroError)
			}
			if divide != 0.0{
				t.Errorf("Division result (%g) should not be equal want: %g.", divide,  0.0 )
			}
		}
	}
}