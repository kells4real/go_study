package main

import (
	"fmt"
	cal "go_study/calculator" // import own function
	test "go_study/speed"     // Another own function import
)

func main() {
	test.Test() // Speed test
	// cal.Calculator() // generic calculator
	fmt.Println(cal.Fibonacci(10))

	// Loop through and print elements in an array 
	var aList = []string{"GET", "POST", "PUT", "DELETE"}
	looped(aList)

}

// Basic loop function
func looped(aList []string) {
	for i := 0; i < len(aList); i++{
		fmt.Println(aList[i])
	}
}