package utils

func Substring(input string, start int, length int) (out string) {
	runes := []rune(input)
	return string(runes[start:length])
}
