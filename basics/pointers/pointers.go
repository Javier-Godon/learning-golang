package main

import "fmt"

func main() {

	var pointer = new(int32)
	var regularVariable int32
	fmt.Printf("Pointer points to: %v\n", pointer)
	fmt.Printf("Pointer points to: %v\n", &pointer)
	fmt.Printf("The value of pointer is: %v\n", *pointer)
	fmt.Printf("The value of regular variable is: %v\n", regularVariable)

	*pointer = 10
	fmt.Printf("The value pointer is: %v\n", *pointer)

	//& to get the address instead of the value

	pointer = &regularVariable
	*pointer = 100
	//now both variables change because they point to the same address
	fmt.Printf("The value of pointer is: %v\n", *pointer)
	fmt.Printf("The value of regular variable is: %v\n", regularVariable)

	var k int32 = 2
	regularVariable = k
	fmt.Printf("The value of regular variable is: %v\n", regularVariable)

	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)
	//this is because slices contains pointers to an underline array. This way we are copying pointers ->
	//both variables refer to the same data

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of the thing1 array is: %p\n", &thing1)
	var result [5]float64 = square(thing1)
	fmt.Printf("\nThe result is: %v\n", result)
	fmt.Printf("\nThe value ot thing1 is: %v\n", thing1)
	//memory locations of thing1 and thing2 are different -> They are different

	//Using pointers
	var thing3 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of the thing3 array is: %p\n", &thing3)
	var result2 [5]float64 = squareUsingPointers(&thing3)
	fmt.Printf("\nThe result is: %v\n", result2)
	fmt.Printf("\nThe value ot thing3 is: %v\n", thing3)

}

// makes a copy of thing1 --> doubles the memory usage: potentially we can be using way more memory than we need
func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the thing2 array is: %p\n", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}

// using pointers (like passing values by reference) -->the value of the original array changes
func squareUsingPointers(thing4 *[5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the thing4 array is: %p\n", thing4)
	for i := range thing4 {
		thing4[i] = thing4[i] * thing4[i]
	}
	return *thing4
}
