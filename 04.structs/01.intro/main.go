// We use structs to represent structured data.
// It's often convenient to group different types of variables together.

package main

import "fmt"

func main() {
	test(messageToSend{
		phoneNumber: "949-267-8598",
		message:     "Thank you for signing up.",
	})
	test(messageToSend{
		phoneNumber: "703-577-5161",
		message:     "Love to have you aboard!",
	})
}

type messageToSend struct {
	phoneNumber string
	message     string
}

func test(m messageToSend) {
	fmt.Printf("Sending message... '%s'\nTo: %v\n", m.message, m.phoneNumber)
	fmt.Println("=======================================")
}
