package piscine

/*
Write a function that simulates the behavior of the Atoi function in Go. Atoi transforms a number defined as a string into a number defined as an int.

Atoi returns 0 if the string is not considered as a valid number. For this exercise only valid string will be tested. They will only contain one or several digits as characters.

For this exercise the handling of the signs + or - does not have to be taken into account.

This function will only have to return the int. For this exercise the error return of Atoi is not required.
*/
func BasicAtoi(s string) int {
	var result int
	// "123"
	for i, ch := range s {
		convChar := int(ch - '0')
		result += convChar
		if i != len(s)-1 {
			result *= 10
		}

	}
	return result
}
