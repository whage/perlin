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

func lerp(a, b, t float64) float64 {
	return a + (b-a)*t
}

func CreatePPM() {
	fmt.Println("P3")
	fmt.Printf("%d %d\n", width, height)
	fmt.Println("255")

	tl := Vec2D{0,4}
	tr := Vec2D{1,-3}
	bl := Vec2D{-2,2}
	br := Vec2D{-3,1}

	var min float64
	var max float64

	pixels := make([][]float64, height)
	for i := 0; i < height; i++ {
		pixels[i] = make([]float64, width)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			value := Fill(tl, tr, bl, br, width, height, x, y)
			pixels[x][y] = value

			if y == 0 && x == 0 {
				min, max = value, value
			} else {
				if value < min {
					min = value
				}
				if value > max {
					max = value
				}
			}
		}
	}

	scaleFactor := 255 / (max-min)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			rgbValue := int((pixels[x][y] - min)*scaleFactor)
			fmt.Printf("%d %d %d\n", rgbValue, rgbValue, rgbValue)
		}
	}
}

func Fill(tl, tr, bl, br Vec2D, width, height, x, y int) float64 {
	v := Vec2D{float64(x), float64(y)}
	dotA := tl.dot(v)
	dotB := tr.dot(v)
	dotC := bl.dot(v)
	dotD := br.dot(v)

	lerpTop := lerp(dotA, dotB, float64(x+1)/float64(width))
	lerpBottom := lerp(dotC, dotD, float64(x+1)/float64(width))

	lerped := lerp(lerpTop, lerpBottom, float64(y+1)/float64(height))
	return lerped
}

/*
func Fill(tl, tr, bl, br Vec2D, width, height int) [][]float64 {
	res := make([][]float64, width*height) // TODO: fix!

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			v := Vec2D{float64(width), float64(height)}
			dotA := tl.dot(v)
			dotB := tr.dot(v)
			dotC := bl.dot(v)
			dotD := br.dot(v)

			lerpTop := lerp(dotA, dotB, float64(width)/float64(x))
			lerpBottom := lerp(dotC, dotD, float64(width)/float64(x))

			res[x][y] = lerp(lerpTop, lerpBottom, float64(height)/float64(y))
		}
	}

	return res
}
*/
