package Itoa

func Itoa(n int) string {
	var (
		reverse = ""
		result = ""
	)

	for n > 0 {
		digit := n%10
		reverse += string(digit+'0')
		n/=10
	}

	for i:= len(reverse)-1; i>=0; i--{
		result += string(reverse[i])
	}

	return result
}