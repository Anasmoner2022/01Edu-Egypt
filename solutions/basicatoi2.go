package piscine

func BasicAtoi2(s string) int {
	var result int

	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			result = result*10 + int(ch-'0')
		} else {
			return 0
		}
	}
	return result
}
