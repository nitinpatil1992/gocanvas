package main

import "fmt"

type Rectangle struct {
	height int
	width  int
}

func (r Rectangle) draw() {
	if r.height == r.width {
		fmt.Println("Drawing square:")
	} else {
		fmt.Println("Drawing rectangle:")
	}
	for i := 0; i < r.height; i++ {
		for j := 0; j < r.width; j++ {
			if i == 0 || j == 0 || i == (r.height-1) || j == (r.width-1) {
				fmt.Printf("%s", "X")
			} else {
				fmt.Printf("%s", " ")
			}
		}
		fmt.Println("")
	}
}
