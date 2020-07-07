package main

import (
	"fmt"
	"math"
)

var i float64
var tres float64
var cinco float64

func main() {
	for i = 0; i >= 100; i++ {
		tres = math.Mod(i, 3)
		cinco = math.Mod(i, 5)

		if tres != 0 && cinco != 0 {
			fmt.Sprintln(i)
		} else if tres == 0 && cinco != 0 {
			fmt.Println("Fizz")
		} else if tres != 0 && cinco == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println("FizzBuzz")
		}
	}
}
