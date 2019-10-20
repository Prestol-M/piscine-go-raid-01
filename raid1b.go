package student

import (
	"github.com/01-edu/z01"
)

func Raid1a(x, y int) {
	if x > 0 {
		for row := 0; row < y; row++ {
			for column := 0; column < x; column++ {
				if row == 0 && column == 0 || column == x-1 && row == y-1 && column > 0 && row > 0 {
					z01.PrintRune('/')
				} else if column == x-1 && row == 0 || column == 0 && row == y-1 {
					z01.PrintRune(92)
				} else if row == 0 || row == y-1 || row <= y-2 && column == 0 || row <= y-2 && column == x-1 {
					z01.PrintRune('*')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
