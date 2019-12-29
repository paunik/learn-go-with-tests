package iteration

const repeatCount = 5

// Repeat repeats string
func Repeat(what string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += what
	}

	return repeated
}
