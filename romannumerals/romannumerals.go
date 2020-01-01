package romannumerals

import "strings"

// RomanNumeral represents roman numeral and arabic value
type RomanNumeral struct {
	Value  int
	Symbol string
}

var RomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// ConvertToRoman converts arabic digit into Roman numeral
func ConvertToRoman(a int) string {
	var result strings.Builder
	for _, numeral := range RomanNumerals {
		for a >= numeral.Value {
			result.WriteString(numeral.Symbol)
			a -= numeral.Value
		}
	}

	return result.String()
}
