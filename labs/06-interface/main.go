package main

import (
	"fmt"
	"math"
)

// Interface - `Shape`
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Struct type `Rectangle` - yg akan implement interface `Shape`
type Rectangle struct {
	Length, Width float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

// Struct type `Circle` - yg akan implement interface `Shape`
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Diameter() float64 {
	return 2 * c.Radius
}

// Generic Function : Bisa menghitung luas area, untuk shape yg berbeda2
func CalculateTotalArea(shapes ...Shape) float64 {
	totalArea := 0.0
	for _, s := range shapes {
		totalArea += s.Area()
	}
	return totalArea
}

// Interfaces can also be used as fields
type MyDrawing struct {
	shapes  []Shape
	bgColor string
	fgColor string
}

func (drawing MyDrawing) Area() float64 {
	totalArea := 0.0
	for _, s := range drawing.shapes {
		totalArea += s.Area()
	}
	return totalArea
}

func main() {
	var s Shape = Circle{5.0}
	fmt.Printf("Type Shape = %T, Nilai Shape %v\n", s, s)
	fmt.Printf("Area = %f, Perimeter %f\n", s.Area(), s.Perimeter())

	var c Circle = Circle{5}
	fmt.Println(c.Diameter())

	s = Rectangle{6.0, 4.0}
	fmt.Printf("Type Shape = %T, Nilai Shape %v\n", s, s)
	fmt.Printf("Area = %f, Perimeter %f\n\n", s.Area(), s.Perimeter())

	totalArea := CalculateTotalArea(Circle{2}, Rectangle{4, 3}, Rectangle{6, 3})
	fmt.Println("Total Area = ", totalArea)

	drawing := MyDrawing{
		shapes: []Shape{
			Circle{2},
			Rectangle{3, 5},
			Rectangle{4, 7},
		},
		bgColor: "red",
		fgColor: "white",
	}

	fmt.Println("Drawing", drawing)
	fmt.Println("Drawing Area = ", drawing.Area())
}
