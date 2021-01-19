package main

import "fmt"
import . "ch6/geometry"

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q))
}
