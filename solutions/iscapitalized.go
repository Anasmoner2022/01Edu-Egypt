package piscine

func IsCapitalized(s string) bool {
	words := Fields(s)
	if len(words) == 0 {
		return false
	}

	for _, word := range words {
		firstChar := word[0]

		if firstChar >= 'a' && firstChar <= 'z' {
			return false
		}
	}
	return true
}
