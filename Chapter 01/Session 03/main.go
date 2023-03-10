package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	fmt.Println("This is your last sentence that you made: ")
	fmt.Println(sentence)
	fmt.Println()

	arrayString := []string{
		" ", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R",
		"S", "T", "U", "V", "W", "X", "Y", "Z",
	}

	lenght := len(arrayString)

	arrayInt := make([]int, lenght)

	for i := 0; i < lenght; i++ {
		arrayInt[i] = strings.Count(sentence, strings.ToLower(arrayString[i]))
	}

	printer(arrayString, arrayInt)

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

func printer(arrayString []string, arrayInt []int) {
	var lenght int
	if len(arrayInt) == len(arrayString) {
		lenght = len(arrayInt)
	} else {
		log.Fatal("Error")
	}

	finalArray := []string{}
	var output string
	for i := 0; i < lenght; i++ {
		if arrayInt[i] != 0 { 
			number := arrayInt[i]
			finalNumber := strconv.Itoa(number)
			output = arrayString[i] + ":" + finalNumber
			finalArray = append(finalArray, output)
		}
	}

	fmt.Printf("map%v\n", finalArray)
}