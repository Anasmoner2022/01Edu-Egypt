package piscine

import "strconv"

func FromTo(from int, to int) string {
	if from > 99 || from < 0 || to > 99 || to < 0 {
		return "Invalid\n"
	}
	var result string
	if from < to {
		for i := from; i <= to; i++ {
			if i > from {
				result += ", "
			}
			if i < 10 {
				result += "0"
			}
			result += strconv.Itoa(i)
		}
	} else {
		for i := from; i >= to; i-- {
			if i < from {
				result += ", "
			}
			if i < 10 {
				result += "0"
			}
			result += strconv.Itoa(i)
		}
	}
	return result + "\n"
}
