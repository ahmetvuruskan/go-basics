package main

import (
	"fmt"
	"strings"
)

func main() {
	//	name := "İnanç"
	//fmt.Println(len(name))
	//fmt.Println(utf8.RuneCountInString(name))
	//yellItBackExample(os.Args[1], "!")
	//yellItBackExercise(os.Args[1], "!")
	const (
		monday = iota + 1
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	const (
		EST = -(5 + iota)
		_
		MST
		PST
	)
}

func yellItBackExample(message, icon string) (int, string) {

	l := len(message)

	s := message + strings.Repeat(icon, l)

	s = strings.ToUpper(s)
	return l, s
}

func yellItBackExercise(message, icon string) {
	l, msg := yellItBackExample(message, icon)

	msg = strings.Repeat(icon, l) + msg

	fmt.Println(msg)
}
