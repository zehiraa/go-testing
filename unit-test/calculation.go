package main

func main(){
	Sum(15.0,8.0)
	Divide(15.0,8.0)
}

func Sum(x float64, y float64) float64 {
	return x + y
}
func Divide(x float64, y float64) (float64, error) {
	if y == 0{
		return 0 , DivideByZeroError
	}
	return x/y, nil
}
