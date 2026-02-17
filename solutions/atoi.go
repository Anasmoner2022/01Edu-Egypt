package piscine

/*
Write a function that simulates the behavior of the Atoi function in Go. Atoi transforms a number represented as a string in a number represented as an int.

Atoi returns 0 if the string is not considered as a valid number. For this exercise non-valid string chains will be tested. Some will contain non-digits characters.

For this exercise the handling of the signs + or - does have to be taken into account.

This function will only have to return the int. For this exercise the error result of Atoi is not required.
*/
func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	var result int
	sign := 1
	start := 0

	switch s[0] {
	case '+':
		sign = 1
		start = 1
	case '-':
		sign = -1
		start = 1
	}

	for i := start; i < len(s); i++ {
		ch := s[i]
		if ch < '0' || ch > '9' {
			return 0
		}
		result = result*10 + int(ch-'0')
	}
	return result * sign
}
