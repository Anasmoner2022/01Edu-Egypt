package piscine

func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Input\n"
	}

	// Remove spaces
	cleaned := ""
	for _, ch := range str {
		if ch != ' ' {
			cleaned += string(ch)
		}
	}
	// fmt.Println(cleaned)
	if len(cleaned) < 5 {
		return "Invalid Input\n"
	}

	result := ""
	position := 0

	for i, ch := range cleaned {
		position++

		if position == 6 {
			position = 0
			continue
		}

		if position <= 4 {
			result += string(ch)
		} else if position == 5 {
			if i != len(cleaned)-1 {
				result += string(ch) + " "
			} else {
				result += string(ch)
			}
		}
	}

	return result + "\n"
}
