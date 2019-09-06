package chess

type Point struct {
	X, Y int
}

func NewPoint(X, Y int) *Point {
	ret := new(Point)
	ret.X = X
	ret.Y = Y
	return ret
}
