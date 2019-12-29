package fizzbuzz

import (
	"errors"
	"strconv"
)

// Fizz string
const Fizz = "Fizz"

// Buzz string
const Buzz = "Buzz"

// ErrInvalidInput error
var ErrInvalidInput = errors.New("Invalid input")

func divisible(numerator, denominator int) bool {
	return numerator%denominator == 0
}

// FizzBuzz implementation
func FizzBuzz(n int) (string, error) {
	var res string

	if n < 0 {
		return "", ErrInvalidInput
	}
	if n == 1 {
		return "1", nil
	}

	for i := 1; i <= n; i++ {
		fizz := divisible(i, 3)
		buzz := divisible(i, 5)

		switch true {
		case fizz && buzz:
			res += Fizz + Buzz
		case fizz:
			res += Fizz
		case buzz:
			res += Buzz
		default:
			res += strconv.Itoa(i)
		}
	}

	return res, nil
}
