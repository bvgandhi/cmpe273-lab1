package main

import (
	"math"
	"testing"
)

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

//TestSleep to test userSleep
func TestPerimeter(t *testing.T) {
	c := userCircle{0, 0, 5}
	if Round(c.area(), .5, 2) != 78.54 {
		t.Errorf("Area of circle is incorrect: %.2f \n", c.area())
	}

	if Round(c.perimeter(), .5, 2) != 31.42 {
		t.Errorf("Perimeter of circle is incorrect: %.2f \n", c.perimeter())
	}

	c = userCircle{0, 0, 21}
	if Round(c.perimeter(), .5, 2) != 131.95 {
		t.Errorf("Perimeter of circle is incorrect: %.2f \n", c.perimeter())
	}

	r := userRectangle{0, 0, 10, 10}
	if Round(r.perimeter(), .5, 2) != 40 {
		t.Errorf("Perimeter of rectangle is incorrect: %.2f \n", r.perimeter())
	}

	r = userRectangle{0, 0, 6, 4}
	if Round(r.perimeter(), .5, 2) != 20 {
		t.Errorf("Perimeter of rectangle is incorrect: %.2f \n", r.perimeter())
	}

}
