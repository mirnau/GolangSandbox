package main

import (
	m "example/hello/mathlib"
	"fmt"
)

func myMessage(hello string, world string) string {

	return hello + " " + world
}

// Comments are C-like;

/*
	Also like this
*/

const PI float64 = 3.14

func main() {

	var output string = myMessage("Hello", "World")
	fmt.Println(output)

	output = "New Value"
	fmt.Println(output)

	fmt.Printf("output: %v\n", output)

	rest := " added and is printed"

	output += rest

	fmt.Println(output)

	var test_array = [3]uint8{12, 64, 121}

	for i := 0; i < int(len(test_array)); i++ {

		fmt.Println(test_array[i])
	}

	// List like type called slice

	var slices = []string{"knife", "cake", "cheese"}

	for i, value := range slices {
		fmt.Println(i, value)
	}

	v := m.Vec3_64{X: 1.0, Y: 2.0, Z: 3.0}

	fmt.Println(v)
}
