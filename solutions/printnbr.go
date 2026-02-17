package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		if n == -9223372036854775808 {
			z01.PrintRune('9')
			n = -223372036854775808
		}
		n = -n
	}

	digits := []rune{}
	if n == 0 {
		digits = append(digits, '0')
	} else {
		for n > 0 {
			digits = append(digits, rune('0'+n%10))
			n /= 10
		}
	}

	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}
