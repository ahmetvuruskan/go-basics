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
	if feetAsInteger, err := strconv.Atoi(feet); err == nil {
		meter := float64(feetAsInteger) * 0.3048

		fmt.Printf("%d feet is %g meter", feetAsInteger, meter)
	} else {
		fmt.Printf("ERROR : %q is not a number  \n", feet)
	}

}
