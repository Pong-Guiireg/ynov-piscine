package piscine

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 2
	}
	for !IsPrime(nb) {
		nb++
	}
	return nb
}
