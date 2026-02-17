package piscine

func Unmatch(a []int) int {
	if len(a) == 0 {
		return -1
	}

	temp := make([]int, len(a))
	copy(temp, a)

	sortSlice(temp)

	for i := 0; i < len(temp)-1; i += 2 {
		if temp[i] != temp[i+1] {
			return temp[i]
		}
	}

	if len(temp)%2 != 0 {
		return temp[len(temp)-1]
	}

	return -1
}

func sortSlice(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
