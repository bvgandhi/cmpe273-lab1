package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type userCircle struct {
	x, y, r float64
}

func circleArea(c *userCircle) float64 {
	return math.Pi * c.r * c.r
}

func (c *userCircle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *userCircle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

type userRectangle struct {
	x1, y1, x2, y2 float64
}

func (r *userRectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *userRectangle) perimeter() float64 {

	return 2 * (distance(r.x1, r.y1, r.x1, r.y2) + distance(r.x1, r.y1, r.x2, r.y1))
}

type userShape interface {
	area() float64
	perimeter() float64
}

func main() {
	fmt.Println("Entering Main")
	c := userCircle{0, 0, 5}
	fmt.Println(circleArea(&c))

	fmt.Println(c.area())
	fmt.Println(c.perimeter())

	r := userRectangle{0, 0, 10, 10}
	fmt.Println(r.area())
	fmt.Println(r.perimeter())
	fmt.Println(math.Pi)
	fmt.Println("Exiting Main")

}
