package main

import (
	"fmt"
)

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	positive, negative := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(positive(i), negative(-2*i))
	}
}
