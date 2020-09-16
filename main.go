package main

import (
	"fmt"
	"go-testing/calculation"
	"go-testing/greeter"
)

func main() {
	calculation.Sum(15.0, 8.0)
	calculation.Divide(15.0, 8.0)

	greeterReader := greeter.NewReader()
	greeterService := greeter.NewGreeterService(greeterReader)

	greeterMessage := greeterService.Greet("en")
	fmt.Println(greeterMessage)

	greeterMessageDef := greeterService.GreetWithDefaultMessage()
	fmt.Println(greeterMessageDef)
}
