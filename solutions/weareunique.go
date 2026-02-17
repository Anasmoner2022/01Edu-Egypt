package piscine

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	map1 := make(map[rune]bool)
	map2 := make(map[rune]bool)

	for _, ch := range str1 {
		map1[ch] = true
	}

	for _, ch := range str2 {
		map2[ch] = true
	}

	count := 0

	for ch := range map1 {
		if !map2[ch] {
			count++
		}
	}

	for ch := range map2 {
		if !map1[ch] {
			count++
		}
	}

	return count
}

// func main() {
// 	fmt.Println(piscine.WeAreUnique("", ""))
// 	fmt.Println(piscine.WeAreUnique("foo", "boo"))
// 	fmt.Println(piscine.WeAreUnique("abc", "def"))
// 	fmt.Println(piscine.WeAreUnique("hello", "world"))   // 5
// 	fmt.Println(piscine.WeAreUnique("aaa", "aaa"))       // 0
// 	fmt.Println(piscine.WeAreUnique("abc", "abc"))       // 0
// 	fmt.Println(piscine.WeAreUnique("a", ""))            // 1
// 	fmt.Println(piscine.WeAreUnique("", "b"))            // 1
// 	fmt.Println(piscine.WeAreUnique("abcdef", "ghijkl")) // 12
// }
