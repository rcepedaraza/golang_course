package main

import "fmt"

// It's critical in Textio that we keep track of how many SMS messages we send on behalf of our clients.
// Fix the bug to accurately track the number of SMS messages sent.
// 1. Alter the incrementSends function so that it returns the result after
// incrementing the sendsSoFar.
// 2. Alter main () to capture the return value from incrementSends ( )
// and overwrite the previous sendsSoFar value.

func main() {
	sendsSoFar := 430
	const sendsToAdd = 25
	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("You've sent: ", sendsSoFar, " messages.")
}

func incrementSends(soFar, toAdd int) int {
	soFar = soFar + toAdd
	return soFar
}
