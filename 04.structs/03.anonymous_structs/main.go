// An anonymous struct is just like a normal struct, but it is defined without a name
// and therefore cannot be referenced elsewhere in the code.

// To create an anonymous struct, just instantiate the instance immediately
// using a second pair of brackets after declaring the type.

myCar := struct {
	make string
	model string
}{
	make: "Tesla"
	model: "model_3"
}

// nesting anonymous structs as fields within other structs.

type car struct {
	Make string
	Model string
	Height int
	Weight int
	// Wheel is a field containing an anonymous struct
	Wheel struct {
		Radius int
		Material string
	}
}