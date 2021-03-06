package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width float64
	height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64{
	return r.width * r.height
}
func (r rect) perim() float64{
	return (r.width + r.height) * 2
}

func (c circle) area() float64{
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func factory(g geometry)  {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 4,height: 3}
	c := circle{radius: 4}

	factory(r)
	factory(c)
}