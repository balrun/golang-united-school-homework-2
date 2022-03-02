package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type Sides int

const (
	SidesCircle   Sides = 0
	SidesTriangle Sides = 3
	SidesSquare   Sides = 4
)

func CalcSquare(sideLen float64, sidesNum Sides) float64 {
	var square float64 = 0
	if sidesNum == SidesTriangle {
		square = sideLen * math.Sqrt(3) / 4
	} else if sidesNum == SidesSquare {
		square = math.Pow(sideLen, 2)
	} else if sidesNum == SidesCircle {
		square = math.Pi * math.Pow(sideLen, 2)
	}
	return square
}
