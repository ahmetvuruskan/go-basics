package main

import (
	"fmt"
	"path"
)

func main() {
	//makeItBlue()
	//varsToVars()
	//assignWithExpressions()
	//calculateRectanglePerimeter()
	//multiAssign()
	//multiAssignSecond()
	//multiShort();
	//swapper();
	//secondSwapper()
	discardFile()
}

func makeItBlue() {
	color := "green"
	color = "blue"

	fmt.Println(color)
}

func varsToVars() {
	color := "green"
	color = "dark " + color
	fmt.Println(color)
}

func assignWithExpressions() {
	n := 0.

	n = 2 * 3.14

	fmt.Println(n)
}

func calculateRectanglePerimeter() {
	var (
		width  = 6
		height = 5
	)
	perimeter := 2 * (width + height)

	fmt.Println(perimeter)
}

func multiAssign() {
	var (
		lang    string
		version int
	)
	lang, version = "go", 2

	fmt.Println(lang, "version", version)
}

func multiAssignSecond() {
	var (
		planet string
		isTrue bool
		temp   float64
	)
	planet = "Mars"
	isTrue = true
	temp = 19.5

	fmt.Println("Air is good on ", planet)
	fmt.Println("It's ", isTrue)
	fmt.Println("It is ", temp, "degrees")

}

func multiShort() {
	_, b := multi()

	fmt.Println(b)
}
func multi() (int, int) {
	return 5, 4
}

func swapper() {
	firstColor, secondColor := "red", "blue"

	firstColor, secondColor = "orange", "green"

	fmt.Println(firstColor, secondColor)
}

func secondSwapper() {
	red, blue := "red", "blue"

	blue, red = red, blue

	fmt.Println(blue, red)
}

func discardFile() {

	folder, _ := path.Split("secret/file.txt")

	fmt.Println(folder)

}
