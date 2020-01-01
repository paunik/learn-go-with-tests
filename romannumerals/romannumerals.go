package romannumerals

import "strings"

// RomanNumeral represents roman numeral and arabic value
type RomanNumeral struct {
	Value  int
	Symbol string
}

// RomanNumerals array of RomanNumeral
type RomanNumerals []RomanNumeral

var rn = RomanNumerals{
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
	for _, numeral := range rn {
		for a >= numeral.Value {
			result.WriteString(numeral.Symbol)
			a -= numeral.Value
		}
	}

	return result.String()
}

// ValueOf method gets value of basic RomanNumerals
func (rn RomanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)
	for _, s := range rn {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

// ConvertToArabic converts from Roman numeral to arabic number
func ConvertToArabic(r string) int {
	total := 0

	for i := 0; i < len(r); i++ {
		symbol := r[i]

		if couldBeSubtractive(i, symbol, r) {
			nextSymbol := r[i+1]
			value := rn.ValueOf(symbol, nextSymbol)

			if value != 0 {
				total += value
				i++
			} else {
				total += rn.ValueOf(symbol)
			}
		} else {
			total += rn.ValueOf(symbol)
		}
	}
	return total
}

func couldBeSubtractive(index int, currentSymbol uint8, roman string) bool {
	isSubtractiveSymbol := currentSymbol == 'I' || currentSymbol == 'X' || currentSymbol == 'C'
	return index+1 < len(roman) && isSubtractiveSymbol
}
