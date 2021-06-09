package reverse

import "unicode/utf8"

func Reverse(s string) string{

	rune := []rune(s)
	var reverse string = ""

	for i := utf8.RuneCountInString(s) - 1; i >= 0; i-- {
		reverse += string(rune[i])
	}
	return string(reverse)
}