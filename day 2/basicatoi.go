package piscine

func BasicAtoi(s string) int {
	num := 0
	for _, char := range s {
		num = num*10 + int(char-'0')
	}
	return num
}
