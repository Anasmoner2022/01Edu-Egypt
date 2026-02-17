package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n < 0 || n > 10 {
		return
	}

	comp := make([]int, n)
	for i := range comp {
		comp[i] = i
	}

	for {
		for _, digit := range comp {
			z01.PrintRune(rune('0' + digit))
		}
		i := n - 1
		for i >= 0 {
			if comp[i] == 0 {
			}
			comp[i]++
		}
		i--
		if i < 0 {
			break
		}
	}
}
