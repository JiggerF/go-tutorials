package main

import "fmt"

type Animal interface {
	Sound() string
	Name() string
}

type Person struct {
	Alias  string
	Gender string
}

func main() {
	p := Person{"Jigger", "F"}
	fmt.Println(p.Sound(), p.Name())
}

func (p Person) Sound() string {
	if p.Alias != "Jigger" {
		return "Woof unknown"
	}
	return "Woof Jigger"
}

func (p Person) Name() string {
	if p.Gender != "M" {
		return "Hello maam!"
	}
	return "Hello Sir!"
}
