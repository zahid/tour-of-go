package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
// 1, 1, 2, 3, 5, 8, 13
func fibonacci() func() int {
	current := 1
	previous := 0
	return func() int {
		temp := current
		current += previous
		previous = temp
		return current
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
