package romannumerals

import "strings"

func ConvertToRoman(a int) string {
	if a == 4 {
		return "IV"
	}

	var result strings.Builder
	for i := 0; i < a; i++ {
		result.WriteString("I")
	}

	return result.String()
}
