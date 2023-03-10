package service

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	color "github.com/TwiN/go-color"
)

func InputInteger(sentence string) int {
	fmt.Println(color.Colorize(color.Bold, color.Ize(color.Blue, "__$ " + sentence)))
	fmt.Print(">>> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	output, _ := strconv.Atoi(input)

	return output
}

func InputString(sentence string) string {
	fmt.Println(color.Colorize(color.Bold, color.Ize(color.Blue, "__$ " + sentence)))
	fmt.Print(">>> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	io := scanner.Text()

	return io
}

func InputCharacter(sentence string) string {
	fmt.Println(color.Colorize(color.Bold, color.Ize(color.Blue, "__$ " + sentence)))
	fmt.Print(">>> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	io := scanner.Text()

	return io
}

func InputYesOrNo(sentence string) bool {
	fmt.Println(color.Colorize(color.Bold, color.Ize(color.Blue, "__$ " + sentence)))
	fmt.Print(">>> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	if strings.ToLower(input) == "y" {
		return true
	} else {
		return false
	}
}
