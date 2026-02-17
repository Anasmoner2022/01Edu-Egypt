package piscine

func HashCode(dec string) string {
	size := len(dec)
	result := ""

	for i := 0; i < size; i++ {
		hashed := (int(dec[i]) + size) % 127

		if hashed < 33 || hashed > 127 {
			hashed += 33
		}
		result += string(rune(hashed))
	}

	return result
}

// func HashCode(dec string) string {
// 	size := len(dec)
// 	result := ""

// 	for _, ch := range dec {
// 		hashed := (int(ch) + size) % 127

// 		if hashed < 33 || hashed > 127 {
// 			hashed += 33
// 		}
// 		result += string(rune(hashed))
// 	}

// 	return result
// }

// func HashCode(dec string) string {
// 	size := len(dec)
// 	result := make([]byte, size)

// 	for i := 0; i < size; i++ {
// 		hashed := (int(dec[i]) + size) % 127

// 		if hashed < 33 || hashed > 127 {
// 			hashed += 33
// 		}
// 		result[i] = byte(hashed)
// 	}

// 	return string(result)
// }

// func HashCode(dec string) string {
// 	size := len(dec)
// 	result := make([]rune, size)

// 	for i := 0; i < size; i++ {
// 		hashed := (int(dec[i]) + size) % 127

// 		if hashed < 33 || hashed > 127 {
// 			hashed += 33
// 		}
// 		result[i] = rune(hashed)
// 	}

// 	return string(result)
// }

// func HashCode(dec string) string {
// 	size := len(dec)
// 	result := ""

// 	for i := 0; i < size; i++ {
// 		hashed := (int(dec[i]) + size) % 127
// 		hashed = makePrintable(hashed)
// 		result += string(rune(hashed))
// 	}

// 	return result
// }

// func makePrintable(ascii int) int {
// 	if ascii < 33 || ascii > 127 {
// 		return ascii + 33
// 	}
// 	return ascii
// }
