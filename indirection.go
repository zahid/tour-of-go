package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(5)
	ScaleFunc(&v, 1.5)

	p := &Vertex{1, 2}
	p.Scale(3)
	ScaleFunc(p, 10)

	fmt.Println(v, p)
}
