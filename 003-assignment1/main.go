package main

import (
	"fmt"
	"math"
)

func main() {

	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, myInt := range ints {
		if math.Mod(float64(myInt), 2) == 0 {
			fmt.Printf("%v is even\n", myInt)
		} else {
			fmt.Printf("%v is odd\n", myInt)
		}

	}

}
