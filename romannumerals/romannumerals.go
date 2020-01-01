package romannumerals

import "strings"

func ConvertToRoman(a int) string {
	var result strings.Builder
	for i := a; i > 0; i-- {
		if i == 4 {
			result.WriteString("IV")
			break
		}
		result.WriteString("I")
	}

	return result.String()
}
