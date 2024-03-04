// Embedded structs provide a kind of data-only inheritance that can be useful at times.
// Go doesn't support classes or inheritance in the complete sense,
// embedded structs are just a way to elevate and share fields between struc definitions.

// type car struct {
// 	make string
// 	model string
// }

// type truck struct {
// 	car
// 	bedSize int
// }

// lanesTruck := truck{
// 	bedSize: 10,
// 	car: car {
// 		make: "Toyota",
// 		model: "Camry"
// 	}
// }

// fmt.println(lanesTruck.bedSize, lanesTruck.make)

package main

import "fmt"

type sender struct {
	rateLimit int
	user
}

type user struct {
	name   string
	number int
}

// don't edit below this line

func test(s sender) {
	fmt.Println("Sender name:", s.name)
	fmt.Println("Sender number:", s.number)
	fmt.Println("Sender rateLimit:", s.rateLimit)
	fmt.Println("====================================")
}

func main() {
	test(sender{
		rateLimit: 10000,
		user: user{
			name:   "Deborah",
			number: 18055558790,
		},
	})
	test(sender{
		rateLimit: 5000,
		user: user{
			name:   "Sarah",
			number: 19055558790,
		},
	})
	test(sender{
		rateLimit: 1000,
		user: user{
			name:   "Sally",
			number: 19055558790,
		},
	})
}
