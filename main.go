package main

import (
	"fmt"
	cal "go_study/calculator" // import own function
	test "go_study/speed"     // Another own function import
)

func main() {
	test.Test() // Speed test
	// cal.Calculator() // generic calculator
	fmt.Println(cal.Fibonacci(15))
}