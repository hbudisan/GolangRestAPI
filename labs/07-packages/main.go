package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(math.Max(73.15, 92.46))
	fmt.Println(math.Sqrt(225))
	fmt.Println(math.Pi)
	fmt.Println(time.Now().Unix())
	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Intn(100))
}
