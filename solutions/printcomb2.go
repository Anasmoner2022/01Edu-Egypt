package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for a1 := '0'; a1 <= '9'; a1++ {
		for a2 := '0'; a2 <= '9'; a2++ {
			for b1 := a1; b1 <= '9'; b1++ {
				for b2 := '0'; b2 <= '9'; b2++ {
					if a1 < b1 || (a1 == b1 && a2 < b2) {
						z01.PrintRune(a1)
						z01.PrintRune(a2)
						z01.PrintRune(' ')
						z01.PrintRune(b1)
						z01.PrintRune(b2)
						if !(a1 == '9' && a2 == '8' && b1 == '9' && b2 == '9') {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
