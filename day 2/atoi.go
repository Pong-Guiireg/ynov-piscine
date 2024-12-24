package piscine

func Atoi(s string) int {
	sign, num := parseSign(&s)
	if num == -1 {
		return 0
	}
	return num * sign
}

func parseSign(s *string) (int, int) {
	sign, num := 1, 0
	previousChar := ' '

	for i := 0; i < len(*s); i++ {
		char := (*s)[i]
		if char == '-' {
			if previousChar == '-' || previousChar == '+' {
				return 0, -1
			}
			sign = -1
		} else if char == '+' {
			if previousChar == '-' || previousChar == '+' {
				return 0, -1
			}
		} else if char >= '0' && char <= '9' {
			num = num*10 + int(char-'0')
		} else {
			return 0, -1
		}
		previousChar = rune(char)
	}
	return sign, num
}
