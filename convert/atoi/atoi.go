package atoi

import (
	"errors"
	"strconv"
)

func Atoi(s string) (int,error) {
	// check to empty
	if s == "" {
		return 0, errors.New("empty name provided")
	}
	_, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("wrong name provided")
	}

	var reverse string
	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	var result int = 0
	bs := []byte(reverse)
	n := 1
	for i := 0; i < len(reverse); i++ {
		result += int(bs[i]-48) * n
		n *= 10
	}
	return result, nil
}