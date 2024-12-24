package piscine

import "github.com/01-edu/z01"

// Vérifie si c'est la dernière combinaison pour n chiffres
func isLastCombination(current []int, n int) bool {
	for i := 0; i < n; i++ {
		if current[i] != 9-(n-1-i) {
			return false
		}
	}
	return true
}

// Imprime une seule combinaison
func printCombination(current []int, n int) {
	for i := 0; i < n; i++ {
		z01.PrintRune(rune(current[i] + '0'))
	}
}

// Imprime les séparateurs (virgule et espace)
func printSeparator() {
	z01.PrintRune(',')
	z01.PrintRune(' ')
}

// Génère les combinaisons de manière récursive
func generateCombinations(current []int, pos, start, n int) {
	if pos == n {
		printCombination(current, n)
		if !isLastCombination(current, n) {
			printSeparator()
		}
		return
	}
	
	for i := start; i <= 9; i++ {
		current[pos] = i
		generateCombinations(current, pos+1, i+1, n)
	}
}

func PrintCombN(n int) {
	current := make([]int, n)
	generateCombinations(current, 0, 0, n)
	z01.PrintRune('\n')
}

func main() {
	PrintCombN(1)
	PrintCombN(3)
	PrintCombN(9)
}

