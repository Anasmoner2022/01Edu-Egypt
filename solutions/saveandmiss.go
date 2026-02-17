// start to code here
package piscine

func SaveAndMiss(arg string, num int) string {
	if num <= 0 {
		return arg
	}
	result := ""
	for i := 0; i < len(arg); i = num*2 + i {
		if i+num < len(arg) {
			result += arg[i : num+i]
		} else {
			result += arg[i:]
		}
	}
	return result
}
