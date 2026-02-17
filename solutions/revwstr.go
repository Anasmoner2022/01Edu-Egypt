package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func RevWStr() {
	if len(os.Args) != 2 {
		return
	}
	sentance := os.Args[1]
	fields := Fields(sentance)
	for i := len(fields) - 1; i >= 0; i-- {
		field := fields[i]
		for _, ch := range field {
			z01.PrintRune(ch)
		}
		z01.PrintRune(' ')
	}
	z01.PrintRune('\n')
}
