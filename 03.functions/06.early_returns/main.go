// Go supports the ability to return early from a function.
// This is a powerful feature that can clean up code, especially when used as guard clauses.
package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := divide(8, 0)
	fmt.Println("Result: ", result)
	fmt.Println("Error: ", err)

}

func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	return dividend / divisor, nil
}
