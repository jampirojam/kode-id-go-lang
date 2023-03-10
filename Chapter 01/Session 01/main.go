package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Print Integer and it data type
	i := 21
	fmt.Printf("%d \n%T \n\n", i /* For numbers */, i /* For data types */)

	// Print special character
	fmt.Printf("%% \n\n")

	// Print boolean value
	j := true
	fmt.Printf("%t \n\n", j)

	// Print Go Escape to unicode value
	unicode := '\u042f'
	fmt.Printf("%c \n\n", unicode)

	// Print final result after convert Integer value of number 21 to Octal and Decimal
	base8 := 21
	base10 := 21
	fmt.Printf("%d \n%o\n\n", base8, base10)

	// Print final result after convert Integer value to Hexadecimal
	var base16 int64 = 15
	valueBase16 := strconv.FormatInt(base16, 16)
	fmt.Printf("%s or %s \n\n", valueBase16, strings.ToUpper(valueBase16))

	base16 = 3859
	valueBase16 = strconv.FormatInt(base16, 16)
	fmt.Printf("%s or %s \n\n", valueBase16, strings.ToUpper(valueBase16))

	// Print index value of an unicode
	unicode = 'Я'
	fmt.Printf("%U \n\n", unicode)

	// Print number of zero after comma or in scientific at float number
	var k float64 = 123.456
	fmt.Printf("%.6f \n%e", k, k)
}
