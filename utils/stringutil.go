package utils

import "strings"

func SubstringBetween(str string, open string, close string) string {
	start := strings.Index(str, open)
	if start != -1 {
		end := strings.Index(str, close)
		if end != -1 {
			return str[start+len(open) : end]
		}
	}
	return ""
}

func RemoveStart(str string, remove string) string {
	start := strings.Index(str, remove)
	if start != -1 {
		return str[start+len(remove):]
	}
	return str
}
