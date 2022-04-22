package main

import "fmt"

func main() {
	f0001println()
	f0002()
	f0003()
}

// Print out some text
func f0001println() {
	println("Go go!")
}

// Create two variables and add them together
func f0002() {
	var a int = 8
	var b int = 34
	var c int = a + b
	println(fmt.Sprintf("%d + %d = %d", a, b, c))
}

// Create an array loop through it and print out the sum of its elements
func f0003() {
	a := [4]int{1, 2, 3, 4}
	b := 0

	// CAVEAT: the underscore (_) represents the index
	//         if we want the index we can use a variable
	//         name for it. but it is required else 'num'
	//         is the index
	for _, num := range a {
		b += num
	}

	println(fmt.Sprintf("sum num %d", b))
}
