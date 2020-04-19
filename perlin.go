package perlin

import (
	"fmt"
	"math/rand"
	"strings"
)

type Vec2D struct {
	X, Y float64
}

func (v Vec2D) dot(o Vec2D) float64 {
	return v.X * o.X + v.Y * o.Y
}

func lerp(a, b, t float64) float64 {
	return a + (b-a)*t
}

func fillGridCell(tl, tr, bl, br Vec2D, width, height int) [][]int {
	var min float64
	var max float64

	values := make([][]float64, height)
	greyScalePixels := make([][]int, height)

	for i := 0; i < height; i++ {
		values[i] = make([]float64, width)
		greyScalePixels[i] = make([]int, width)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			value := fillPoint(tl, tr, bl, br, width, height, x, y)
			values[x][y] = value

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
			greyScalePixels[x][y] = int((values[x][y] - min)*scaleFactor)
		}
	}

	return greyScalePixels
}

func CreatePPM(width, height int) {
	fmt.Println("P3")
	fmt.Printf("%d %d\n", width, height)
	fmt.Println("255")

	gridWidth := 4
	gridHeight := 3

	gridNodeVectors := make([][]Vec2D, gridHeight+1)

	for j := 0; j < gridHeight; j++ {
		gridNodeVectors[j] = make([]Vec2D, gridWidth+1)
		for i := 0; i < gridWidth; i++ {
			gridNodeVectors[j][i] = Vec2D{rand.Float64(),rand.Float64()}
		}
	}

	pixelValues := make([][]string, height)
	for j := 0; j < height; j++ {
		pixelValues[j] = make([]string, width)
	}

	for y := 0; y < gridHeight-1; y++ {
		for x := 0; x < gridWidth-1; x++ {
			tl := gridNodeVectors[y][x]
			tr := gridNodeVectors[y][x+1]
			bl := gridNodeVectors[y+1][x]
			br := gridNodeVectors[y+1][x+1]

			cellWidth := width/gridWidth
			cellHeight := height/gridHeight

			greyScaleValues := fillGridCell(tl, tr, bl, br, cellWidth, cellHeight)
			
			for n := 0; n < cellHeight; n++ {
				for m := 0; m < cellWidth; m++ {
					pixelValues[x*cellWidth+m][y*cellHeight+n] = fmt.Sprintf("%d %d %d", greyScaleValues[m][n], greyScaleValues[m][n], greyScaleValues[m][n])
				}
			}
		}
	}

	for j := 0; j < height; j++ {
		var row []string
		for i := 0; i < width; i++ {
			row = append(row, pixelValues[j][i])
		}
		fmt.Println(strings.Join(row, " "))
	}
}

func fillPoint(tl, tr, bl, br Vec2D, width, height, x, y int) float64 {
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
