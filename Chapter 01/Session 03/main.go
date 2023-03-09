package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var sentence string
	menu(sentence)
}

func menu(sentence string) {
	fmt.Println("Write random sentences below, double enter to see the result.")
	fmt.Println("To make blank on the line, use space then enter to continue writing")
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()
		if len(line) != 0 {
			sentence = sentence + scanner.Text()
		} else {
			break
		}
	}

	counter(strings.ToLower(sentence))
}

func counter(sentence string) {
	countA := strings.Count(sentence, "a")
	countB := strings.Count(sentence, "b")
	countC := strings.Count(sentence, "c")
	countD := strings.Count(sentence, "d")
	countE := strings.Count(sentence, "e")
	countF := strings.Count(sentence, "f")
	countG := strings.Count(sentence, "g")
	countH := strings.Count(sentence, "h")
	countI := strings.Count(sentence, "i")
	countJ := strings.Count(sentence, "j")
	countK := strings.Count(sentence, "k")
	countL := strings.Count(sentence, "l")
	countM := strings.Count(sentence, "m")
	countN := strings.Count(sentence, "n")
	countO := strings.Count(sentence, "o")
	countP := strings.Count(sentence, "p")
	countQ := strings.Count(sentence, "q")
	countR := strings.Count(sentence, "r")
	countS := strings.Count(sentence, "s")
	countT := strings.Count(sentence, "t")
	countU := strings.Count(sentence, "u")
	countV := strings.Count(sentence, "v")
	countW := strings.Count(sentence, "w")
	countX := strings.Count(sentence, "x")
	countY := strings.Count(sentence, "y")
	countZ := strings.Count(sentence, "z")
	countSpace := strings.Count(sentence, " ")

	// Print using Multi Array
	arrayMulti := [][]interface{}{
		{" ", countSpace}, {"A", countA}, {"B", countB}, {"C", countC}, {"D", countD}, {"E", countE},
		{"F", countF}, {"G", countG}, {"H", countH}, {"I", countI}, {"J", countJ}, {"K", countK},
		{"L", countL}, {"M", countM}, {"N", countN}, {"O", countO}, {"P", countP}, {"Q", countQ},
		{"R", countR}, {"S", countS}, {"T", countT}, {"U", countU}, {"V", countV}, {"W", countW},
		{"X", countX}, {"Y", countY}, {"Z", countZ},
	}

	printerMulti(arrayMulti)

	// Print using Single Array
	arrayInt := []int{
		countSpace, countA, countB, countC, countD, countE, countF, countG, countH, countI, countJ,
		countK, countL, countM, countN, countO, countP, countQ, countR, countS, countT, countU,
		countV, countW, countX, countY, countZ,
	}

	arrayString := []string{
		" ", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R",
		"S", "T", "U", "V", "W", "X", "Y", "Z",
	}

	printSingle(arrayString, arrayInt)

	fmt.Println("This is your last sentence that you made: ")
	fmt.Println(sentence)
	fmt.Println()
	fmt.Println("Continue counting with more sentence [Y/N]")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice := scanner.Text()

	if strings.ToLower(choice) == "y" {
		fmt.Println()
		menu(sentence + ". ")
	} else {
		log.Fatal("Error")
		os.Exit(0)
	}
}

func printerMulti(arrayMulti [][]interface{}) {
	var i, j int
	fmt.Println("PRINT USING MULTI ARRAY")
	for i = 0; i < len(arrayMulti); i++ {
		for j = 0; j < len(arrayMulti[0]); j++ {
			if j%2 == 0 {
				fmt.Printf("%v:", arrayMulti[i][j])
			} else {
				fmt.Printf("%v ", arrayMulti[i][j])
			}
		}
	}
	fmt.Printf("\n\n")
}

func printSingle(arrayString []string, arrayInt []int) {
	var lenght int
	if len(arrayInt) == len(arrayString) {
		lenght = len(arrayInt)
	} else {
		log.Fatal("Error")
	}

	fmt.Println("PRINT USING SINGLE ARRAY")
	for i := 0; i < lenght; i++ {
		if i == (lenght - 1) {
			fmt.Printf("%v:%v", arrayString[i], arrayInt[i])
		} else {
			fmt.Printf("%v:%v ", arrayString[i], arrayInt[i])
		}
	}
	fmt.Printf("\n\n")
}
