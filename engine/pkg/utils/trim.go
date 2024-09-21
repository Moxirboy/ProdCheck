package utils
import "strings"

func SplitBeforeAfterFirstSlash(input string) (string, string) {
	firstSlashIndex := strings.Index(input, "/")
	if firstSlashIndex != -1 {
		before := input[:firstSlashIndex]
		after := input[firstSlashIndex+1:] // +1 to skip the '/'
		return before, after
	}
	return input, "" // If no '/' is found, return the input and an empty string
}