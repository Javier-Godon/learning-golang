package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "specialize"
	var indexed = myString[0]
	fmt.Println(indexed)
	fmt.Printf("%v, %T", indexed, indexed)
	fmt.Println()

	var myString2 = "résumë"
	for index, value := range myString2 {
		fmt.Println(index, value)
	}

	//a string is treated as an array of strings -> the length is represented in strings
	fmt.Printf("\nThe length of the string is %v", len(myString2))

	//run is an alias to int35. This way I can get the length of the string in characters
	var myString3 = []rune("résumë")
	fmt.Printf("\nThe length of the string is %v", len(myString3))

	//We can declare a run type with single quotes
	var myRune = 'a'
	fmt.Printf("\nmyRune is = %v", myRune)
	//String concatenation with + symbol
	var strSlice = []string{"u", "n", "d", "e", "r"}
	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Printf("\n%v", catStr)
	fmt.Println()

	//To manipulate strings we can use strings ->strings.Builder
	var stringBuilder = strings.Builder{}
	for i := range strSlice {
		stringBuilder.WriteString(strSlice[i])
	}
	var builtString = stringBuilder.String()
	fmt.Printf("\n%v with stringBuilder", builtString)
	fmt.Println()

}
