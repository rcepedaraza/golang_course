package main

import "fmt"

// We often will need to manipulate strings in our messaging app.
// For example, adding some personalization by using a customer's name within a template.
// The concat function should take two strings and smash them together.

func concat(string1 string, string2 string) string { // function signature
	return string1 + string2
}

func main() {
	fmt.Println(concat("Ricardo, ", "happy birthday!"))
	fmt.Println("Eva, ", "hope that Tesla stock works!")
}
