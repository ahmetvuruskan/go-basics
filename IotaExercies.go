package main

import "fmt"

func main() {
	months()
	monthsSecond()
	seasons()
}

func months() {
	const (
		Nov = 11 - iota
		Oct
		Sep
	)

	fmt.Println(Sep, Oct, Nov)

}

func monthsSecond() {

	const (
		_ = iota
		Jan
		Feb
		Mar
	)

	// This should print: 1 2 3
	// Not: 0 1 2
	fmt.Println(Jan, Feb, Mar)
}

func seasons() {

	const (
		Spring = (1 + iota) * 3
		Summer
		Fall
		Winter
	)

	fmt.Println(Winter, Spring, Summer, Fall)
}
