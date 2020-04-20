package perlin

import (
	"fmt"
	"math"
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

func fillGridCell(tl, tr, bl, br Vec2D, cellWidth, cellHeight int) [][]int {
	var min float64
	var max float64

	values := make([][]float64, cellWidth)
	greyScalePixels := make([][]int, cellWidth)

	for i := 0; i < cellWidth; i++ {
		values[i] = make([]float64, cellHeight)
		greyScalePixels[i] = make([]int, cellHeight)
	}

	for y := 0; y < cellHeight; y++ {
		for x := 0; x < cellWidth; x++ {
			value := fillPoint(tl, tr, bl, br, cellWidth, cellHeight, x, y)
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

	for y := 0; y < cellHeight; y++ {
		for x := 0; x < cellWidth; x++ {
			greyScalePixels[x][y] = scaleToRGBRange(values[x][y], min, max)
		}
	}

	return greyScalePixels
}

func scaleToRGBRange(value, min, max float64) int {
	scaleFactor := 255 / (max-min)
	return int((value-min) * scaleFactor)
}

func CreatePPM(width, height, gridWidth, gridHeight int) {
	fmt.Println("P3")
	fmt.Printf("%d %d\n", width, height)
	fmt.Println("255")

	gridNodeVectors := make([][]Vec2D, gridWidth+1)

	for i := 0; i <= gridWidth; i++ {
		gridNodeVectors[i] = make([]Vec2D, gridHeight+1)
		for j := 0; j <= gridHeight; j++ {
			angle := rand.Float64() * math.Pi
			gridNodeVectors[i][j] = Vec2D{math.Cos(angle),math.Sin(angle)}
		}
	}

	pixelValues := make([][]string, width)
	for i := 0; i < width; i++ {
		pixelValues[i] = make([]string, height)
	}

	for y := 0; y < gridHeight; y++ {
		for x := 0; x < gridWidth; x++ {
			tl := gridNodeVectors[x][y]
			tr := gridNodeVectors[x+1][y]
			bl := gridNodeVectors[x][y+1]
			br := gridNodeVectors[x+1][y+1]

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
			row = append(row, pixelValues[i][j])
		}
		fmt.Println(strings.Join(row, " "))
	}
}

func fillPoint(tl, tr, bl, br Vec2D, width, height, x, y int) float64 {
	dotA := tl.dot(Vec2D{float64(x), float64(y)})
	dotB := tr.dot(Vec2D{float64(x-width), float64(y)})
	dotC := bl.dot(Vec2D{float64(x), float64(y-height)})
	dotD := br.dot(Vec2D{float64(x-width), float64(y-height)})

	lerpTop := lerp(dotA, dotB, float64(x)/float64(width))
	lerpBottom := lerp(dotC, dotD, float64(x)/float64(width))

	lerped := lerp(lerpTop, lerpBottom, float64(y)/float64(height))
	return lerped
}
