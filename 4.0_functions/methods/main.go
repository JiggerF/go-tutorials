package main

import "fmt"

type Square struct {
	length, width, height int
}

func main() {
	s := Square{12, 1, 7}
	fmt.Println("Total area: ", s.Area())
}

// Using Methods
func (s Square) Area() int {
	return s.length * s.width * s.height
}
