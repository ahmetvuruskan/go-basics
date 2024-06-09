package main

import "fmt"

func main() {
	//first()
	//second()
	//third()
	//fourth()
	fifth()
}

// ---------------------------------------------------------
// EXERCISE: Convert and Fix
//
//	Fix the code by using the conversion expression.
//
// a, b := 10, 5.5
// fmt.Println(a + b)
// EXPECTED OUTPUT
//
//	15.5
//
// ---------------------------------------------------------
func first() {
	a, b := 10, 5.5

	addition := float64(a) + b

	fmt.Print(addition)
}

// ---------------------------------------------------------
// EXERCISE: Convert and Fix #2
//
//  Fix the code by using the conversion expression.
// a, b := 10, 5.5
// a = b
// fmt.Println(a + b)
// EXPECTED OUTPUT
//  10.5
// ---------------------------------------------------------

func second() {
	a, b := 10, 5.5
	a = int(b)
	fmt.Println(a)
	fmt.Println(float64(a) + b)
}

// ---------------------------------------------------------
// EXERCISE: Convert and Fix #3
//
//	Fix the code.
//
// EXPECTED OUTPUT
// fmt.Println(int(5.5))
//
//	5.5
//
// ---------------------------------------------------------
func third() {
	fmt.Println(5.5)
}

// ---------------------------------------------------------
// EXERCISE: Convert and Fix #4
//
//	Fix the code.
//
// EXPECTED OUTPUT
//
//	9.5
//
// ---------------------------------------------------------
func fourth() {

	age := 2
	fmt.Println(7.5 + float64(age))
}

// ---------------------------------------------------------
// EXERCISE: Convert and Fix #5
//
//	Fix the code.
//
// HINTS
//
//	maximum of int8  can be 127
//	maximum of int16 can be 32767
//
// EXPECTED OUTPUT
//
//	1127
//
// ---------------------------------------------------------
func fifth() {

	minV := int8(127)
	maxV := int16(1000)

	fmt.Println(maxV + int16(minV))
}
