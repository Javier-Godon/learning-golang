package main

func updateName(name string) {
	name = "Elias"
}

func updateNamePassingPointer(name *string) {
	*name = "Elias"
}

func main() {
	name := "John"
	updateName(name)
	println(name)
	updateNamePassingPointer(&name)
	println(name)

	var otherName = new(string)
	*otherName = "Sam"
	updateNamePassingPointer(otherName)
	println(*otherName)

}
