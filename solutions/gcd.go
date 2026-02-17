package piscine

func Gcd(a, b int) int {
	if a == 0 && b == 0 {
		return 0
	}

	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// func Gcd(a, b int) int {
// 	if a == 0 && b == 0 {
// 		return 0
// 	}

// 	if b == 0 {
// 		return a
// 	}

// 	return Gcd(b, a%b)
// }
