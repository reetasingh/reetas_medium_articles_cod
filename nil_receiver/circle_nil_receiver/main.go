package main

import (
	"fmt"
	"math"
)

type Circle struct {
	r float64
}

func (c *Circle) area() float64 {
	if c == nil {
		return 0
	}
	return math.Pi * c.r * c.r
}

func main() {
	var c *Circle
	fmt.Printf("(%v, %T)\n", c, c)
	fmt.Println(c.area())
}
