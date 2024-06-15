package main

import (
	"fmt"
	"os"
)

func main() {
	//countArguments()
	//printPath()
	//printYourName()
	//greetMorePeople()
}

// ---------------------------------------------------------
// EXERCISE: Count the Arguments
//
//  Print the count of the command-line arguments
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 names.
// ---------------------------------------------------------

func countArguments() {
	count := len(os.Args) - 1

	fmt.Printf("There are %d names \n", count)
}

// ---------------------------------------------------------
// EXERCISE: Print the Path
//
//	Print the path of the running program by getting it
//	from `os.Args` variable.
//
// HINT
//
//	Use `go build` to build your program.
//	Then run it using the compiled executable program file.
//
// EXPECTED OUTPUT SHOULD INCLUDE THIS
//
//	myprogram
//
// ---------------------------------------------------------
func printPath() {
	fmt.Printf(os.Args[0])
}

func printYourName() {
	var name string = os.Args[1]

	fmt.Println(name + "\n")
	fmt.Printf("Hi ! %s \n", name)
	fmt.Println("How are you")
}

func greetMorePeople() {
	var length = len(os.Args) - 1
	if length < 3 {
		fmt.Println("You should type minimum three name you typed ", length)
		os.Exit(length)
	}
	fmt.Printf("There are %d peoples \n", length)

	for i := 1; i <= length; i++ {
		fmt.Printf("Hello great %s \n", os.Args[i])
	}
}
