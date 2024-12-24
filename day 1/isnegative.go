package main

import (
	"github.com/01-edu/z01"
)

func IsNegative(n int) {
	if n < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
}
