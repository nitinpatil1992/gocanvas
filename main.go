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

		case 3:

		case 4:
			os.Exit(0)
		}
	}

}
