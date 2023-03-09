package main

import (
	"fmt"
)

func main() {
	var i int
	var j int
	var unicode string = "САШАРВО"

	for i = 0; i <= 4; i++ {
		fmt.Printf("Nilai i = %d\n", i)
	}

	for j = 0; j <= 10; j++ {
		if (j == 5) {
			for number, char := range unicode {
				fmt.Printf("Character %#U starts at byte position %d\n", char, number)
			}
		} else {
			fmt.Printf("Nilai j = %d\n", j)
		}
	}
}