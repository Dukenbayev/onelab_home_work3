package runebyindex

func RuneByIndex(s *string, i *int) (r rune, err error) {
	defer func() {
		if r := recover(); r != nil {
			if wrapped, ok := r.(error); ok {
				err = wrapped
			}
		}
	}()
	runes := []rune(*s)
	return runes[*i], nil
}