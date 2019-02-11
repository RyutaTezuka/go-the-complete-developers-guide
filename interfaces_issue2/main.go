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
	sq := square{
		sideLength: 10.0,
	}

	tr := triangle{
		height: 5.0,
		base:   10.0,
	}
	printArea(sq)
	printArea(tr)
}

func printArea(sp shape) {
	fmt.Println(sp.getArea())
}
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}
