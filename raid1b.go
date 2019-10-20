package student

import (
	"github.com/01-edu/z01"
)

func Raid1b(x, y int) {
	if x > 0 {
		for row := 0; row < y; row++ {
			if row == 0 || row == y-1 {
				for column := 0; column < x; column++ {
					if column == 0 && row == 0 || column == x-1 && row == y-1 && column > 1 && row > 1 {
						z01.PrintRune('/')
					} else if column == x-1 && row == 0 || column == 0 && row == y-1 {
						z01.PrintRune(92)
					} else {
						z01.PrintRune('*')
					}
				}
				z01.PrintRune('\n')
			} else if row != 0 && row != y-1 {
				for column := 0; column < x; column++ {
					if column == 0 || column == x-1 {
						z01.PrintRune('*')
					} else {
						z01.PrintRune(' ')
					}
				}
				z01.PrintRune('\n')
			}

		}
	}
}
