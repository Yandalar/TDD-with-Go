package iteration

func Repeat(character string, numberOfIterations int) string {
	var repeated string
	for i := 0; i < numberOfIterations; i++ {
		repeated += character
	}

	return repeated
}
