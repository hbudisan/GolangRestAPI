package main

import "fmt"

// Struct Type : Point
type Point struct {
	X, Y float64
}

// Method with Receiver `Point`
func (p Point) IsAbove(y float64) bool {
	return p.Y > y
}

// Method diubah jadi function
func IsAbove(p Point, y float64) bool {
	return p.Y > y
}

func main() {
	p := Point{2.0, 4.0}
	fmt.Println("Point ", p)

	fmt.Println("Apakah Point p terletak diatas garis y = 1.0 ", p.IsAbove(1.0))
	fmt.Println("Apakah Point p terletak diatas garis y = 1.0 ", IsAbove(p, 1.0))
}
