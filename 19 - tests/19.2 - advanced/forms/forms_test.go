package forms

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		r := Rectangle{10, 10}
		expectedArea := float64(100)
		receivedArea := r.Area()

		if expectedArea != receivedArea {
			t.Errorf("The Aread expected %f is different from received %f", expectedArea, receivedArea)
		}

	})

	t.Run("Circle", func(t *testing.T) {
		c := Circle{10}
		expectedArea := float64( math.Pi * 100)
		receivedArea := c.Area()

		if expectedArea != receivedArea {
			t.Errorf("The Aread expected %f is different from received %f", expectedArea, receivedArea)
		}
	})
}
