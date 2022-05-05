package main

import "fmt"

type bot interface {
	calculateArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type rectangle struct {
	sidelength float64
}

func main() {
	tx := triangle{2, 4}
	rx := rectangle{3}
	printArea(tx)
	printArea(rx)

}

func printArea(b bot) {
	fmt.Println(b.calculateArea())
}

func (t triangle) calculateArea() float64 {
	return t.height * t.base * 0.5
}

func (r rectangle) calculateArea() float64 {
	return r.sidelength * r.sidelength
}

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type geometry interface {
// 	area() float64
// }

// type rect struct {
// 	width  float64
// 	height float64
// }
// type circle struct {
// 	radius float64
// }

// func (r rect) area() float64 {
// 	return r.width * r.height
// }

// func (c circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// func measure(g geometry) {
// 	fmt.Println(g)
// 	fmt.Println(g.area())
// }

// func main() {
// 	r := rect{width: 3, height: 4}
// 	c := circle{radius: 5}

// 	measure(r)
// 	measure(c)
// }
