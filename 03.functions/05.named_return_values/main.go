// Return values may be fiven names, and if they are, then they are treated the same
// as if they were new variables defined at the top of the function.
// Named return values are best thought of as a way to document the purpose of the retuned values.

package main

import "fmt"

func main() {
	test(4)
	test(10)
	test(22)
	test(35)
}

func yearsUnitlEvents(age int) (yearsUntilAdult int, yearsUntilDrinking int, yearsUnitlCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUnitlCarRental = 25 - age
	if yearsUnitlCarRental < 0 {
		yearsUnitlCarRental = 0
	}
	return
}

func test(age int) {
	fmt.Println("Age: ", age)
	yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUnitlEvents(age)
	fmt.Println("You are an adult in: ", yearsUntilAdult, " years")
	fmt.Println("You can drink in: ", yearsUntilDrinking, " years")
	fmt.Println("You can rent a car in: ", yearsUntilCarRental, " years")
	fmt.Println("==========================================")
}
