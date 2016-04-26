package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

func describe(i I) {
	fmt.Printf("%v, %T\n", i, i)
}
