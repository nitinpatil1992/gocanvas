package main

import "fmt"

type Triangle struct {
	side int
}

func (s Triangle) draw() {
	fmt.Println("Drawing traingle:")
	x := s.side - 1

	for i := 0; i < s.side; i++ {
		for j := 0; j <= 2*x; j++ {
			if j == (x-i) || j == (x+i) || (i == x && j%2 == 0) {
				fmt.Printf("%s", "X")
			} else {
				fmt.Printf("%s", " ")
			}
		}
		fmt.Println("")
	}
}
