package main

import "fmt"

func main() {
	//Arrays: Fixed Length, same type, Indexable, Contiguous in Memory
	var intArr [3]int32
	fmt.Println(intArr)
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])
	intArr[0] = 1
	intArr[1] = 2
	intArr[2] = 3
	fmt.Println(intArr[1:3])
	fmt.Println(intArr[0])
	fmt.Println(intArr[1])
	fmt.Println(intArr[2])
	//They are stored in contiguous memory locations
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	secondIntArr := [3]int64{4, 5, 6}
	fmt.Println(secondIntArr[0])
	fmt.Println(secondIntArr[1])
	fmt.Println(secondIntArr[2])
	//They are stored in contiguous memory locations
	fmt.Println(&secondIntArr[0])
	fmt.Println(&secondIntArr[1])
	fmt.Println(&secondIntArr[2])

	thirdIntArr := [...]int64{4, 5, 6}
	fmt.Println(thirdIntArr[0:2])

	fourthIntArr := [...]int64{4, 5, 6}
	fmt.Println(fourthIntArr)

	//Slices wrap arrays to give a more general, powerful, and convenient interface to sequences of data (like List in java)
	var intSlice = []int32{4, 5, 6}
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	fmt.Println("")
	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	fmt.Println("")

	var intSlice2 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	intSlice3 := make([]int32, 3)
	fmt.Println(intSlice3)
	fmt.Printf("The length is %v with capacity %v", len(intSlice3), cap(intSlice))
	fmt.Println("")

	//Maps
	myMap := make(map[string]uint8)
	fmt.Println(myMap)
	myMap["Jonny"] = 35
	myMap["Frank"] = 42
	myMap["Sarah"] = 47
	fmt.Println(myMap)
	delete(myMap, "Jonny")
	fmt.Println(myMap)

	myMap2 := map[string]uint8{"Adam": 23, "John": 42}
	fmt.Println(myMap2)
	fmt.Println(myMap2["Adam"])
	//a map always returns something, even if the key does not exist (it returns the default value, in this case 0)!!!!
	fmt.Println(myMap2["foo"])
	//
	var age, isKeyInTheMap = myMap2["Adam"]
	fmt.Printf("value %v and boolean if the key is in the map %v", age, isKeyInTheMap)
	fmt.Println("")

	age, isKeyInTheMap = myMap2["foo"]
	if isKeyInTheMap {
		fmt.Printf("The age is %v", age)
		fmt.Println("")
	} else {
		fmt.Printf("Invalid name")
		fmt.Println("")
	}

	// min 24:30

}
