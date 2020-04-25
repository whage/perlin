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

func smoothInterpolate(a, b, t float64) float64 {
	rangeWidth := b-a
	if rangeWidth == 0 {
		return a
	}
	//return hermite(t) * rangeWidth + a
	return fifthDegree(t) * rangeWidth + a
}

func hermite(t float64) float64 {
	return 3 * t * t - 2 * t * t * t
}

func fifthDegree(t float64) float64 {
	return 6*math.Pow(t,5) - 15*math.Pow(t, 4) + 10*math.Pow(t, 3)
}

func fillGridCell(tl, tr, bl, br Vec2D, cellWidth, cellHeight int) [][]float64 {
	values := make([][]float64, cellWidth)

	for i := 0; i < cellWidth; i++ {
		values[i] = make([]float64, cellHeight)
	}

	for y := 0; y < cellHeight; y++ {
		for x := 0; x < cellWidth; x++ {
			value := getValueOfPoint(tl, tr, bl, br, cellWidth, cellHeight, x, y)
			values[x][y] = value
		}
	}

	return values
}

func scaleToRGBRange(value, min, max float64) int {
	scaleFactor := 255 / (max-min)
	return int((value-min) * scaleFactor)
}

func generateRandomUnitVectors(gridWidth, gridHeight int) [][]Vec2D {
	gridNodeVectors := make([][]Vec2D, gridWidth+1)

	for i := 0; i <= gridWidth; i++ {
		gridNodeVectors[i] = make([]Vec2D, gridHeight+1)
		for j := 0; j <= gridHeight; j++ {
			angle := rand.Float64() * 2 * math.Pi
			gridNodeVectors[i][j] = Vec2D{math.Cos(angle), math.Sin(angle)}
		}
	}

	return gridNodeVectors
}

func CreatePPM(width, height, gridWidth, gridHeight int) {
	fmt.Println("P3")
	fmt.Printf("%d %d\n", width, height)
	fmt.Println("255")

	var seenFirstValue bool = false
	var min float64
	var max float64

	gridNodeVectors := generateRandomUnitVectors(gridWidth, gridHeight)

	pixelValues := make([][]float64, width)
	for i := 0; i < width; i++ {
		pixelValues[i] = make([]float64, height)
	}

	for y := 0; y < gridHeight; y++ {
		for x := 0; x < gridWidth; x++ {
			tl := gridNodeVectors[x][y]
			tr := gridNodeVectors[x+1][y]
			bl := gridNodeVectors[x][y+1]
			br := gridNodeVectors[x+1][y+1]

			cellWidth := width/gridWidth
			cellHeight := height/gridHeight

			pixelValuesInCell := fillGridCell(tl, tr, bl, br, cellWidth, cellHeight)
			
			for n := 0; n < cellHeight; n++ {
				for m := 0; m < cellWidth; m++ {
					if !seenFirstValue {
						min, max = pixelValuesInCell[m][n], pixelValuesInCell[m][n]
						seenFirstValue = true
					}
					if pixelValuesInCell[m][n] < min {
						min = pixelValuesInCell[m][n]
					}
					if pixelValuesInCell[m][n] > max {
						max = pixelValuesInCell[m][n]
					}
					pixelValues[x*cellWidth+m][y*cellHeight+n] = pixelValuesInCell[m][n]
				}
			}
		}
	}

	for j := 0; j < height; j++ {
		var row []string
		for i := 0; i < width; i++ {
			scaled := scaleToRGBRange(pixelValues[i][j], min, max)
			row = append(row, fmt.Sprintf("%d %d %d", scaled, scaled, scaled))
		}
		fmt.Println(strings.Join(row, " "))
	}
}

func getValueOfPoint(tl, tr, bl, br Vec2D, width, height, x, y int) float64 {
	dotA := tl.dot(Vec2D{float64(x), float64(y)})
	dotB := tr.dot(Vec2D{float64(x-width), float64(y)})
	dotC := bl.dot(Vec2D{float64(x), float64(y-height)})
	dotD := br.dot(Vec2D{float64(x-width), float64(y-height)})

	top := smoothInterpolate(dotA, dotB, float64(x)/float64(width))
	bottom := smoothInterpolate(dotC, dotD, float64(x)/float64(width))
	final := smoothInterpolate(top, bottom, float64(y)/float64(height))

	return final
}
