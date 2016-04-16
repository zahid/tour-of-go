package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
    var picture = make([]uint8, dx)
    for i := 0, i < dy; i++ {
        picture = append(picture make([]uint8, dy))
    }

    // for i := 0; i < len(picture); i++ {
    //     for j := 0; len(picture[i]); j++ {

    //     }
    // }
    return picture
}

func main() {
    pic.Show(Pic)
}

