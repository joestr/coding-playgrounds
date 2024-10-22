package main

import "fmt"

func main() {
	f0001println()
	f0002()
	f0003()
	f0004()
	f0005()
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

// Create a 3x3 matrix, fill in 1 to 9 then print out
func f0004() {
	a := make([][]int, 3)
	for i := 0; i < 3; i++ {
		a[i] = make([]int, 3)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			a[i][j] = (i * 3) + j + 1
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d,", a[i][j])
		}
		fmt.Printf("\n")
	}
}

// Shift bits
func f0005() {
	var firstBits = 0b00001001
	var secondBits = 0b00000011
	var resultBits = firstBits >> secondBits
	fmt.Printf("Shift bits: %d >> %d = %d", firstBits, secondBits, resultBits)
}

func onePlusOne() int {
	return 1 + 1
}
