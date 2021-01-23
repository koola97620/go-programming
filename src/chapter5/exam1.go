package main

import (
	"fmt"
	"math"
)

func hypot(x, y float64) (r float64) {
	r = math.Sqrt(x*x + y*y)
	return r
}

func main() {
	fmt.Println(hypot(3, 4)) // 5
}
