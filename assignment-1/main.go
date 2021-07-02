package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	sq := square{sideLength: 4}
	tr := triangle{base: 3, height: 6}

	printArea(sq)
	printArea(tr)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (tr triangle) getArea() float64 {
	return 0.5 * tr.base * tr.height
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}
