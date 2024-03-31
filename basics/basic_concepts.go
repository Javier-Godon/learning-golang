package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum32 float32 = 12345678.9
	fmt.Println(floatNum32)

	var floatNum64 float64 = 12345678.9
	fmt.Println(floatNum64)

	var result float32 = floatNum32 + float32(floatNum64)
	fmt.Println(result)

	const intNum1 = 3
	const intNum2 = 2
	fmt.Println(intNum1 / intNum2)

	const intNum1Float float64 = 3
	const intNum2Float float64 = 2
	fmt.Println(intNum1Float / intNum2Float)

	var firstString string = "First string" + " concatenated y"
	fmt.Println(firstString)

	fmt.Println("the length of the string is: " + strconv.FormatInt(int64(len(firstString)), 10))

	fmt.Println(len(firstString))
	fmt.Println(utf8.RuneCountInString(firstString))

	fmt.Println("the length of the string is for sure : " + strconv.FormatInt(int64(utf8.RuneCountInString(firstString)), 10))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool
	fmt.Println(myBoolean)

	var myVarTypical = "typical"
	myVarEasy := "easy"

	fmt.Println(myVarTypical)
	fmt.Println(myVarEasy)

	var1, var2, var3 := 1, 2, 3

	fmt.Println(var1, var2, var3)
	fmt.Println([3]int{var1, var2, var3})

	printMe("to print in printMe")
	fmt.Println(intDivision(100, 5))
	fmt.Println(intDivisionWithRemainder(100, 45))
	var divisionResult, remainder int = intDivisionWithRemainder(11, 2)
	fmt.Printf("The result of the integer division is %v with remainder %v", divisionResult, remainder)
	fmt.Println("")
	var divisionResult1, remainder1, err = intDivisionWithRemainderAndError(11, 0)
	if err != nil {
		fmt.Printf(err.Error())
		fmt.Println("")
	} else if remainder1 == 0 {
		fmt.Printf("The result of the integer division is %v", divisionResult1)
	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v", divisionResult1, remainder1)
	}

	divisionResult1, remainder1, err = intDivisionWithRemainderAndError(11, 3)
	//in switch brake is implicit
	switch {
	case err != nil:
		fmt.Printf(err.Error())
		fmt.Println("")
	case remainder1 == 0:
		fmt.Printf("from switch: The result of the integer division is %v", divisionResult1)
	default:
		fmt.Printf("from switch: The result of the integer division is %v with remainder %v", divisionResult1, remainder1)
	}

	switch remainder {
	case 0:
		fmt.Println("The division was exact")
	case 1, 2:
		fmt.Println("The division was close")
	default:
		fmt.Println("The division was not close")
	}

	fmt.Println("")
	makeChecks(1, 2)

}

func makeChecks(firstCheck int, secondCheck int) {
	if firstCheck == 1 && secondCheck == 2 {
		fmt.Println("check passed")
	}
}

func printMe(valueToPrint string) {
	fmt.Println(valueToPrint)
}

func intDivision(numerator int, denominator int) int {
	return numerator / denominator
}

func intDivisionWithRemainder(numerator int, denominator int) (int, int) {
	var result = numerator / denominator
	var remainder = numerator % denominator
	return result, remainder
}

func intDivisionWithRemainderAndError(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result = numerator / denominator
	var remainder = numerator % denominator
	return result, remainder, err
}
