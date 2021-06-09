package fibonacci

func Fibonacci() func() int {
	first := 0
	second := 0
	return func() int{
		if(first == 0){
			first, second = 1,1
			return 0
		}else {
			current, firstc := second, first
			second,first = first + second, firstc
			return current
		}
	}
}
