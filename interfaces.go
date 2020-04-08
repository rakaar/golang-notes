package main

import "fmt"

type geo interface {
	area() int
	perim() int
}

type cir struct {
	radius int
}

type sq struct {
	side int
}

func (c cir) area() int {
	return 3*c.radius*c.radius
}

func (c cir) perim() int {
	return 6*c.radius 
}

func (s sq) area() int {
	return s.side * s.side
}

func (s sq) perim() int {
	return 4*s.side
}

func measure(g geo) {
	fmt.Println(g,g.area())
	fmt.Println(g,g.perim())
}

func main() {
	sample_c := cir{radius: 2}
	sample_s := sq{ side: 200 }

	measure(sample_s)
	measure(sample_c)
	
}
