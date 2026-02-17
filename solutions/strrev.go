package piscine

func StrRev(s string) string {
	bytes := []byte(s)
	for i := 0; i < len(bytes)/2; i++ {
		j := len(bytes) - 1 - i
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}
