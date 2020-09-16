package calculation

import (
	"testing"
)

// TestSub validates the http Get function can download content and handles different status conditions properly.
func TestSubForDivide(t *testing.T) {
	tests := []struct {
		name string
		x float64
		y float64
		result float64
		err error
	}{
		{"nonZero",10.0, 5.0, 2.0, nil},
		{"zero",1.0, 0.0, 0.0, DivideByZeroError},
	}

	t.Log("Given the need to test dividing different variables.")
	{
		// Range over our table but this time, create an anonymous function that takes a testing T  parameter. This is a test function inside a test function.
		// What nice about it is that we are gonna have a new function for each set of data that we have in our table. Therefore, we will end up with 2 different functions here.
		for _, tt := range tests {
			tf := func(t *testing.T) {

				// The only difference here is that we call Parallel function inside each of these individual sub test function.
				t.Parallel() // example:  parallel_test.go

				t.Logf("When divide %g for %s decimal %g", tt.x, tt.name, tt.y)
				{
					res, err := Divide(tt.x, tt.y)
					if err != nil{
						t.Logf("Should division return error for divide by zero, result: %v, expected: %v.", err,  tt.err )
					}else if res != tt.result {
						t.Errorf("Should division result was incorrect, result: %g, expected: %g.", res,   tt.result  )
					}else{
						t.Logf("Should division was correct, result: %g , expected: %g.", res,  tt.result )
					}
				}
			}
			// Once we declare this function, we tell the testing tool to register it as a sub test under the test name.
			t.Run(tt.name, tf)
		}
	}
}