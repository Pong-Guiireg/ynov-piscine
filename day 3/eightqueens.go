package piscine

import "github.com/01-edu/z01"

func EightQueens() {
    solutions := []string{}
    solveQueens([]int{}, &solutions)
    printSolutions(solutions)
}

func solveQueens(queens []int, solutions *[]string) {
    if len(queens) == 8 {
        *solutions = append(*solutions, formatSolution(queens))
        return
    }
    for col := 0; col < 8; col++ {
        if isSafe(queens, col) {
            queens = append(queens, col)
            solveQueens(queens, solutions)
            queens = queens[:len(queens)-1] // backtrack
        }
    }
}

func isSafe(queens []int, col int) bool {
    row := len(queens)
    for r, c := range queens {
        if c == col || r-c == row-col || r+c == row+col {
            return false
        }
    }
    return true
}

func formatSolution(queens []int) string {
    solution := ""
    for _, col := range queens {
        z01.PrintRune(rune(col + '1')) // convert to 1-based index and print directly
    }
    return solution
}

func printSolutions(solutions []string) {
    for _, solution := range solutions {
        for _, char := range solution {
            z01.PrintRune(char)
        }
        z01.PrintRune('\n')
    }
}
