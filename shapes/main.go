package main

import "fmt"

type (
	shape interface {
		getArea() float64
	}

	triangle struct {
		height float64
		base   float64
	}

	square struct {
		sideLength float64
	}
)

func main() {
	t := triangle{height: 10, base: 15}
	s := square{sideLength: 20}

	printArea(t)
	printArea(s)
}

func printArea(sp shape) {
	fmt.Printf("Shape %T has an area %f", sp, sp.getArea())
}

/*func (sp shape) printArea() {
	fmt.Printf("Shape %t has an area %f", s, s.getArea())
}

[go] invalid receiver type shape (shape is an interface type)
printArea func()
*/

func (t triangle) getArea() float64 {
	return .5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
