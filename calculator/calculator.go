package calculator
import (
	"fmt"
	"strings"
)


func Calculator() float64{
	var result float64 = 0
	var Result = &result // Used this just to get a feel of how pointers work.
	var num1 float64
	var num2 float64
	var sign string


	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter Sign, e.g X for multiplication: ")
	fmt.Scanln(&sign)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	*Result = subCal(num1, num2, sign)
	fmt.Println(result) // Prints out the initial result 

	// An infinite loop to continually run the calculator
	for {
		var fNum float64
		var sign string
		fmt.Print("Enter Sign: ")
		fmt.Scanln(&sign)
		var newSign = strings.ToLower(sign) // convert sign input to lower to check sign
		if newSign == "exit" {
			break
		}else{
		fmt.Print("Enter Number: ")
		fmt.Scanln(&fNum)
		*Result = subCal(*Result, fNum, sign)
		fmt.Println(result)
		}
	}
	return *Result
}


func subCal(num float64, fNum float64, sign string) float64 {

	var result float64

	switch sign {
	case "+":
		result = num + fNum
	case "-":
		result = num - fNum
	case "*", "X", "x":
		result = num * fNum
	case "/":
		result = num / fNum
	}
	return result
}