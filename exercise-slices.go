package main

import (
	"fmt"
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var picture = make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		picture[i] = make([]uint8, dx)
	}
	fmt.Println(picture)
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			intValue := ((x + y) / 2)
			//intValue := ((x * y))
			picture[y][x] = uint8(intValue)
		}
	}
	return picture
}

func main() {
	pic.Show(Pic)
}
