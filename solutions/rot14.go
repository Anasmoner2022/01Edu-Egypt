package piscine

func Rot14(s string) string {
	result := ""
	for _, ch := range s {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
			rot14 := ch + 14
			if (rot14 > 'z' && (ch >= 'a' && ch <= 'z')) || (rot14 > 'Z' && (ch >= 'A' && ch <= 'Z')) {
				rot14 -= 26
			}
			result += string(rot14)
		} else {
			result += string(ch)
		}
	}

	return result
}

// func Rot14(s string) string {
// 	runes := []rune(s)

// 	for i, ch := range s {
// 		if ch >= 'a' && ch <= 'z' {
// 			rot14 := (ch-'a'+14)%26 + 'a'
// 			runes[i] = rot14
// 		} else if ch >= 'A' && ch <= 'Z' {
// 			rot14 := (ch-'A'+14)%26 + 'A'
// 			runes[i] = rot14
// 		}
// 	}
// 	return string(runes)
// }
