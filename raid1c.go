package student

import (
	"github.com/01-edu/z01"
)

func Raid1b(x, y int) {
	if x > 0 {
		for row := 0; row < y; row++ {
			for column := 0; column < x; column++ {
				if row == 0 && column == 0 || row == 0 && column == x-1 {
					z01.PrintRune('A')
				} else if column == 0 && row == y-1 || column == x-1 && row == y-1 {
					z01.PrintRune('C')
				} else if row == 0 || row == y-1 || row <= y-2 && column == 0 || row <= y-2 && column == x-1 {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
