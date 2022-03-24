package main

import "fmt"

type rectangle struct {
	height float64
	width  float64
}
type triangle struct {
	height float64
	width  float64
}

type Shape interface {
	getArea() float64
}

func (r rectangle) getArea() float64 {
	return r.width * r.height
}

func (t triangle) getArea() float64 {
	return t.width * t.height / 2
}

func printArea(s Shape) {
	fmt.Printf("Area = %+v\n", s.getArea())
}

func main() {
	r := rectangle{height: 15, width: 20}
	t := triangle{height: 15, width: 20}
	printArea(r)
	printArea(t)
}
