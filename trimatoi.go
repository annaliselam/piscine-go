package piscine

import (
// "github.com/01-edu/z01"
)

// Write a function that transforms a number within a string, in an int.

// For this exercise the handling of the - sign has to be taken into account. If the sign is encountered before any number it should determine the sign of the returned int.

// This function will only return an int. In case of invalid input, the function should return 0.

// Note: There will never be more than one sign by string in the tests.

func TrimAtoi(s string) int {
	runeString := []rune(s)
	sign := 1
	result := 0
	foundDigit := false

	for _, letter := range runeString {
		if !foundDigit {
			if letter == '-' {
				sign = -1
			} else if letter == '+' {
				sign = 1
			}
		}
		if letter >= '0' && letter <= '9' {
			foundDigit = true
			result = 10*result + int(letter-'0')
		}
	}
	return sign * result
}
