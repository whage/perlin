package perlin

import (
	"fmt"
)

var width = 200
var height = 200

type Vec2D struct {
	X, Y float64
}

func (v Vec2D) dot(o Vec2D) float64 {
	return v.X * o.X + v.Y * o.Y
}

func CreatePNG() {
	fmt.Println("P3")
	fmt.Printf("%d %d\n", width, height)
	fmt.Println("255")

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			fmt.Println("255 255 255")
		}
	}
}

func Fill(tl, tr, bl, br Vec2D, width, height int) [][]float64 {
	res := make([width][height]float64, 0) // TODO: fix!

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			v := Vec2D{width, height}
			dotA := tl.dot(v)
			dotB := tr.dot(v)
			dotC := bl.dot(v)
			dotD := br.dot(v)

			lerpTop := lerp(dotA, dotB)
			lerpBottom := lerp(dotC, dotD)

			res[x][y] = 
		}
	}

	return res
}
