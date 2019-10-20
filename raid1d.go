package student

import (
	"github.com/01-edu/z01"
)

func Raid1d(x, y int) {
	if x > 0 {
		for row := 0; row < y; row++ {
			if row == 0 || row == y-1 {
				for column := 0; column < x; column++ {
					if column == 0 && row == 0 || column == 0 && row == y-1 {
						z01.PrintRune('A')
					} else if column > 1 && column == x-1 && row == y-1 || column == x-1 && row == 0 {
						z01.PrintRune('C')
					} else {
						z01.PrintRune('B')
					}
				}
				z01.PrintRune('\n')
			} else if row != 0 && row != y-1 {
				for column := 0; column < x; column++ {
					if column == 0 || column == x-1 {
						z01.PrintRune('B')
					} else {
						z01.PrintRune(' ')
					}
				}
				z01.PrintRune('\n')
			}

		}
	}
}
