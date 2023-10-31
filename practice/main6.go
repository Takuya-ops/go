package main

import (
	"fmt"
	"math" //円周率
)

type circle struct {
	r int
	h int
}

func (c circle) circle_volume() float64 {
	return float64(c.r*c.r*c.h) * math.Pi
}

func main() {
	// circle := circle{1, 3}
	circle := circle{2, 3}
	fmt.Println(circle.circle_volume())
	fmt.Println(circle)

}
