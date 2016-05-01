package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Zazu", 23}
	b := Person{"Britt", 22}

	fmt.Println(a, b)
}
