package main

import (
	"fmt"
	"os"
)

func main() {
	var option int
	var s Shape
	for {
		fmt.Println("Enter options number for drawing")
		fmt.Println("1) Circle")
		fmt.Println("2) Triangle")
		fmt.Println("3) Rectangle  or Square")
		fmt.Println("4) exit ")

		fmt.Scanln(&option)
		switch option {
		case 1:
			var radius int
			fmt.Println("Enter radius")
			fmt.Scanln(&radius)
			s = Circle{radius}
			s.draw()
		case 2:
			var height int
			fmt.Println("Enter height")
			fmt.Scanln(&height)
			s = Triangle{height}
			s.draw()
		case 3:
			var height, width int
			fmt.Println("Enter height and width with space separated value")
			fmt.Scanf("%d %d", &height, &width)
			s = Rectangle{height, width}
			s.draw()
		case 4:
			os.Exit(0)
		}
	}

}
