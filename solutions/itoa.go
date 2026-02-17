package piscine

func Itoa(n int) string {
	nums := []int{}
	var result string
	if n == 0 {
		return "0"
	}

	if n < 0 {
		n = -n
		result += "-"
	}

	for n != 0 {
		nums = append(nums, n%10)
		n = n / 10
	}
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		result += string(rune(num + '0'))
	}
	return result
}
