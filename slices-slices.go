package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 9}
	fmt.Printf("s == %d\n", s)
	fmt.Printf("s[1:4] == %d\n", s[1:4])
	fmt.Printf("s[:3] == %d\n", s[:3])
	fmt.Printf("s[4: == %d\n", s[4:])
}
