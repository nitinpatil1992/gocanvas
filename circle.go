package main

import "fmt"

type Circle struct {
	radius int
}

func (c Circle) draw() {
	fmt.Println("Drawing circle with radius ", c.radius)

	x := c.radius

	// top part
	for i := 0; i <= c.radius; i++ {
		for j := 0; j <= 2*x; j++ {
			if j == (x-i) || j == (x+i) {
				fmt.Printf("%s", "X")
			} else {
				fmt.Printf("%s", " ")
			}
		}
		fmt.Println("")
	}

	// bottom part
	for i := c.radius - 1; i >= 0; i-- {
		for j := 0; j <= 2*x; j++ {
			if j == (x-i) || j == (x+i) {
				fmt.Printf("%s", "X")
			} else {
				fmt.Printf("%s", " ")
			}
		}
		fmt.Println("")
	}

}
