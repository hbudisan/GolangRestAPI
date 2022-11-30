package main

import (
	"fmt"
	"math"
)

// function declaration
func avg(x float64, y float64) float64 {
	return (x + y) / 2
}

// function with multiple return values
func getStockPriceChange(prevPrice, currentPrice float64) (change, percentChange float64) {
	change = currentPrice - prevPrice
	percentChange = (change / prevPrice) * 100
	return
}

func main() {
	fmt.Println("Functions")

	x := 5.75
	y := 6.25

	result := avg(x, y)
	fmt.Printf("Rata2 dari %.2f dan %.2f = %.2f\n", x, y, result)

	prevStockPrice := 925.0
	currentStockPrice := 825.0

	change, percentChange := getStockPriceChange(prevStockPrice, currentStockPrice)

	if change < 0 {
		fmt.Printf("Harga saham turun sebesar Rp %.2f atau sebesar %.2f%%", math.Abs(change), math.Abs(percentChange))
	} else {
		fmt.Printf("Harga saham naik sebesar Rp %.2f atau sebesar %.2f%%", math.Abs(change), math.Abs(percentChange))
	}
}
