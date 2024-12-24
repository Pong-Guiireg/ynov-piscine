package piscine

func BasicAtoi2(s string) int {
	num := 0
	for _, char := range s {
		if char >= '0' && char <= '9' {
			num = num*10 + int(char-'0')
		} else {
			return 0
		}
	}
	return num
}
