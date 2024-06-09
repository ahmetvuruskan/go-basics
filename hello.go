package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", os.Args)

	fmt.Println("Path:", os.Args[0])
	fmt.Println("First:", os.Args[1])
	fmt.Println("Second:", os.Args[2])

	fmt.Println("Count", len(os.Args))
}
