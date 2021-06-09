package atoi

import (
	"fmt"
	"strconv"
)

func Atoi(s string) (int,error) {
	// check to empty
	if s == "" {
		return 0, fmt.Errorf("can not convert string to int: invalid format")
	}
	_, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("can not convert string to int: invalid format")
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