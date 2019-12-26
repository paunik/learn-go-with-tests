package fizzbuzz

import (
	"errors"
	"strconv"
)

var ErrInvalidInput = errors.New("Invalid input")

func FizzBuzz(n int) (string, error) {

	var res string

	if n < 0 {
		return "", ErrInvalidInput
	}
	if n == 1 {
		return "1", nil
	}

	for i := 1; i <= n; i++ {
		switch true {
		case i%5 == 0 && i%3 == 0:
			res += "FizzBuzz"
		case i%3 == 0:
			res += "Fizz"
		case i%5 == 0:
			res += "Buzz"
		default:
			res += strconv.Itoa(i)
		}
	}

	return res, nil
}
