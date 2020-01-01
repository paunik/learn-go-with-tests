package romannumerals

import "strings"

func ConvertToRoman(a int) string {
	var result strings.Builder
	for a > 0 {
		switch {
		case a > 4:
			result.WriteString("V")
			a -= 5
		case a > 3:
			result.WriteString("IV")
			a -= 4
		default:
			result.WriteString("I")
			a--
		}
	}

	return result.String()
}
