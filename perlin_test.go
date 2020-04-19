package perlin

import (
	"testing"
)

func TestVectorDotProduct(t *testing.T) {
	v1 := Vec2D{2, 3}
	v2 := Vec2D{1, 2}
	result := v1.dot(v2)
	expected := float64(8)
	if result != expected {
		t.Errorf("Expected `%v`, got %v", expected, result)
	}
}
