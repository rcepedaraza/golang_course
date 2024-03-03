// A function can return a value that the caller doesn't care about.
// We can explicitly ignore variables by using an underscore.
package main

import "fmt"

func main() {
	firstName, _ := getNames()
	fmt.Println("Welcome: ", firstName)
}

func getNames() (string, string) {
	return "Ricardo", "Cepeda Raza"
}
