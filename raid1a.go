package main

import (
	"github.com/01-edu/z01"
)

func Raid1a(x, y int) {
	if x > 0 {
		for raw := 0; raw < y; raw++ {
			if raw == 0 || raw == y-1 {
				for column := 0; column < x; column++ {
					if column == 0 || column == x-1 {
						z01.PrintRune('o')
					} else {
						z01.PrintRune('-')
					}
				}
				z01.PrintRune('\n')
			} else if raw != 0 && raw != y-1 {
				for column := 0; column < x; column++ {
					if column == 0 || column == x-1 {
						z01.PrintRune('|')
					} else {
						z01.PrintRune(' ')
					}
				}
				z01.PrintRune('\n')
			}

		}
	}
}

func main() {
	Raid1a(45, 5)
}
