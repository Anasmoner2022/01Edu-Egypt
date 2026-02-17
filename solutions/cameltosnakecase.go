package piscine

func CamelToSnakeCase(s string) string {
	result := ""
	if len(s) == 0 {
		return ""
	}

	if !isValid(s) {
		return s
	}

	for i, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			if i > 0 {
				result += "_"
			}
			result += string(ch)
		} else {
			result += string(ch)
		}
	}
	return result
}

func isValid(s string) bool {
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if i == len(s)-1 && ch >= 'A' && ch <= 'Z' {
			return false
		} else if i < len(s) && (ch >= 'A' && ch <= 'Z') && (s[i+1] >= 'A' && s[i+1] <= 'Z') {
			return false
		} else if !((ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')) {
			return false
		}
	}
	return true
}
