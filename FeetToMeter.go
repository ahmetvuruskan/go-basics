package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	feetToMeter(os.Args[1])
}

func feetToMeter(feet string) {
	feetAsInteger, err := strconv.Atoi(feet)

	if err != nil {
		fmt.Printf("ERROR : %q is not a number  \n")
	}
	meter := float64(feetAsInteger) * 0.3048

	fmt.Printf("%d feet is %g meter", feetAsInteger, meter)

}
