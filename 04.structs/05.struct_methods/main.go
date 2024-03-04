// Methods can be defined on structs.
// methods are just functions that have a receiver.
// a receiver is a special parameter that syntatically goes before the name of the function.

// type rect struct {
// 	width int
// 	height int
// }

// // area has a receiver of (r rect)
// func (r rect) area () int {
// 	return r.width * r.height
// }

// r := rect {
// 	width: 5
// 	height: 10
// }

// fmt.println(r.area())

package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (a authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic \n%s: %s", a.username, a.password)
}

func test(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("====================================")
}

func main() {
	test(authenticationInfo{
		username: "Google",
		password: "12345",
	})
	test(authenticationInfo{
		username: "Bing",
		password: "98765",
	})
	test(authenticationInfo{
		username: "DDG",
		password: "76921",
	})
}
