package util

import "strings"

func StringIsEmpty(str string) bool {
	return strings.TrimSpace(str) == ""
}

func StringIsNotEmpty(str string) bool {
	return !StringIsEmpty(str)
}
