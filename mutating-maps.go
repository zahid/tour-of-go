package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["Answer"] = 2
	fmt.Println("m[Answer]= ", m["Answer"])

	delete(m, "Answer")
	fmt.Println(m)

	v, ok := m["Answer"]

	fmt.Printf("Val %d present %v\n", v, ok)

}
