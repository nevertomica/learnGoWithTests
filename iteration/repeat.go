package iteration

func Repeat(character string, times int) string {
	var repeated string

	if times < 0 {
		times = 0
	}

	for i := 0; i < times; i++ {
		repeated += character
	}
	return repeated
}
