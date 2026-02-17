package piscine

func FirstWord(s string) string {
	fields := Fields(s)
	if len(fields) == 0 {
		return "\n"
	}
	return fields[0] + "\n"
}

func Fields(s string) []string {
	var result []string
	i := 0
	for i < len(s) {

		for i < len(s) && isWhiteSpace(s[i]) {
			i++
		}

		start := i
		for i < len(s) && !isWhiteSpace(s[i]) {
			i++
		}

		if i > start {
			result = append(result, s[start:i])
		}
	}

	return result
}

func isWhiteSpace(c byte) bool {
	return c == ' ' || c == '\n' || c == '\t' || c == '\r'
}
