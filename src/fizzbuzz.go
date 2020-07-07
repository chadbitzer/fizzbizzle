package main

import (
	"fmt"
	"math"
	"strconv"
)

var i float64
var tres float64
var cinco float64

func main() {
	for i = 0; i >= 100; i++ {
		tres = math.Mod(i, 3)
		cinco = math.Mod(i, 5)

		strTres := strconv.FormatFloat(tres, 'E', 2, 64)
		fmt.Println(strTres)
		strCinco := strconv.FormatFloat(cinco, 'E', 2, 64)
		fmt.Println(strCinco)
		//		if tres < 1 && cinco < 1 {
		//			fmt.Sprintln(i)
		//		} else if tres == 0 && cinco < 1 {
		//			fmt.Println("Fizz")
		//		} else if tres < 1 && cinco == 0 {
		//			fmt.Println("Buzz")
		//		} else {
		//			fmt.Println("FizzBuzz")
		//		}
	}
}
