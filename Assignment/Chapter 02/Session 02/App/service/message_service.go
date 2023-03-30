package service

import "net/http"

func ErrorMessage(code int, message string) (string, string) {
	return "message", http.StatusText(code) + "-> " + message
}

func SuccessMessage(code int, message string) (string, string) {
	return "message", http.StatusText(code) + "-> " + message
}