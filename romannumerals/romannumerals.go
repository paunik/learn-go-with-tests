package romannumerals

func ConvertToRoman(a int) string {
	switch a {
	case 1:
		return "I"
	case 2:
		return "II"
	case 3:
		return "III"
	default:
		return ""
	}
}
