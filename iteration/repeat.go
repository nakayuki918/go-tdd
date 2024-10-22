package iteration

func Repeat(characters []string) string {
	var repeated string
	for _, character := range characters {
		repeated += character
	}
	return repeated
}
