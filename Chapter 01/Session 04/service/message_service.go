package service

import (
	"fmt"
	color "github.com/TwiN/go-color"
)

func ErrorMessage(message string) {
	fmt.Println(color.Colorize(color.Red, "Err! ") + message)
}

func SuccessMessage(message string) {
	fmt.Println(color.Colorize(color.Green, "Suc! ") + message)
}
