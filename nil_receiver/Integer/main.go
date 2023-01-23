package main

import "fmt"

type Integer int

func (i Integer) add(j Integer) Integer {
	return i + j
}

func main() {
	i := Integer(1)
	j := Integer(2)
	fmt.Println(i.add(j))
}
