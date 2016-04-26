package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

type F float64

func (t T) M() {
	fmt.Println(t.S)
}

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I
	i = &T{"Hellooo"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%T, %v)\n", i, i)
}
