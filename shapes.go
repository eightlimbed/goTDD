package shapes

import (
	"math"
)

// Shape is an interface for any type that has an Area() method
type Shape interface {
	Area() float64
}

// Rectangle is a four-sided shape
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle is a round shape with a radius
type Circle struct {
	Radius float64
}

// Triangle is a pointy shape
type Triangle struct {
	Base   float64
	Height float64
}

// Perimeter returns the sum of all sides of a Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area returns the area of a Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area returns the area of a Circle
func (c Circle) Area() float64 {
	return (c.Radius * c.Radius) * math.Pi
}

// Area returns the area of a Triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}
