package piscine

import (
	"unicode"

	"github.com/01-edu/z01"
)

func PrintMemory(arr [10]byte) {
	for i := 0; i < len(arr); i++ {
		if i > 0 && i%4 == 0 {
			z01.PrintRune('\n')
		} else if i > 0 {
			z01.PrintRune(' ')
		}
		printHex(arr[i])
	}
	z01.PrintRune('\n')

	for i := 0; i < len(arr); i++ {
		if unicode.IsGraphic(rune(arr[i])) {
			z01.PrintRune(rune(arr[i]))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}

func printHex(b byte) {
	hexValues := "0123456789abcdef"
	z01.PrintRune(rune(hexValues[b/16]))
	z01.PrintRune(rune(hexValues[b%16]))
}
