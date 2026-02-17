package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func PigLatin() {
	if len(os.Args) != 2 {
		return
	}
	word := os.Args[1]
	result := pigLatin(word)
	for _, ch := range result {
		z01.PrintRune(ch)
	}
}

func pigLatin(s string) string {
	if len(s) == 0 {
		return ""
	}

	if !hasVowels(s) {
		return "No vowels"
	}

	if isVowel(rune(s[0])) {
		return s + "ay"
	}

	firstVowelIndex := findFirstVowel(s)

	return s[firstVowelIndex:] + s[:firstVowelIndex] + "ay"
}

func hasVowels(s string) bool {
	for _, ch := range s {
		if isVowel(ch) {
			return true
		}
	}
	return false
}

func isVowel(ch rune) bool {
	ch = toLower(ch)
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}

func findFirstVowel(s string) int {
	for i, ch := range s {
		if isVowel(ch) {
			return i
		}
	}
	return -1
}

func toLower(ch rune) rune {
	if ch >= 'A' && ch <= 'Z' {
		return ch + 32
	}
	return ch
}
