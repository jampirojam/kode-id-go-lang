package service

import "strings"

func GetMarriedStatus(marriedCode string) (string, bool) {
	married := strings.ToLower(marriedCode) == "m"

	var marriedStatus string
	switch strings.ToLower(marriedCode) {
	case "m":
		marriedStatus = "Married"
	case "s":
		marriedStatus = "Single"
	case "w":
		marriedStatus = "Widowed"
	case "d":
		marriedStatus = "Divorce"
	}
	
	return marriedStatus, married
}
