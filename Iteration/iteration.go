package iteration

func Repeat(char string, times int) string {
	var repeated string
	for range times {
		repeated += char
	}
	return repeated
}
