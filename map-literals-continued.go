package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Lon float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
	"ZTown":     {40.771194, -73.793185},
}

func main() {
	fmt.Println(m)
}
