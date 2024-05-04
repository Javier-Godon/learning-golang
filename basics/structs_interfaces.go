package main

import (
	"fmt"
)

type engine interface {
	milesLeft() uint8
}

type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
	owner
}

type owner struct {
	name string
}

// /methods of the structs are created as functions
func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

func canMakeIt(e gasEngine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there")
	} else {
		fmt.Println("Need to fuel up first")
	}
}

func canMakeItUsingInterface(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there")
	} else {
		fmt.Println("Need to fuel up first")
	}
}

func main() {
	var myEngine = gasEngine{mpg: 25, gallons: 15, ownerInfo: owner{"John Smith"}}
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo)

	myEngine.mpg = 10
	myEngine.gallons = 12
	myEngine.ownerInfo = owner{"John Brown"}
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo)
	//anonymous structs
	var myEngine2 = struct {
		mpg     uint8
		gallons uint8
		owner
	}{25, 100, owner{"Michael statham"}}
	fmt.Println(myEngine2.mpg, myEngine2.gallons, myEngine2.name)

	fmt.Printf("Total miles left in tank: %v", myEngine.milesLeft())
	fmt.Println()

	canMakeIt(myEngine, 250)
	var myElectricEngine = electricEngine{10, 10, owner{"Jonny"}}
	canMakeItUsingInterface(myElectricEngine, 50)

}
