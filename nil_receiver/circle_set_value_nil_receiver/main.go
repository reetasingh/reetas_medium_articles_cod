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
		c = new(Circle)
		c.r = 10
	}
	return math.Pi * c.r * c.r
}

func main() 
{
	var c *Circle
	fmt.Println(c.area())
	fmt.Println(c.r)
}
