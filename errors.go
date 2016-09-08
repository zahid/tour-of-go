package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("At %v, %v", e.When, e.What)
}

func run() error {
	return &MyError{time.Now(), "Your mom fell through the floor"}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
