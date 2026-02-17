package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	first := os.Args[1]
	second := os.Args[2]

	inSecond := make(map[rune]bool)
	for _, ch := range second {
		inSecond[ch] = true
	}

	added := make(map[rune]bool)
	result := ""

	for _, ch := range first {
		if inSecond[ch] && !added[ch] {
			result += string(ch)
			added[ch] = true
		}
	}

	result += "\n"
	print(result)
}

func print(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
}
