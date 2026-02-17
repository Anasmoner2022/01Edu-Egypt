package piscine

func LastWord(s string) string {
	fields := Fields(s)
	if len(fields) == 0 {
		return "\n"
	}
	return fields[len(fields)-1] + "\n"
}
