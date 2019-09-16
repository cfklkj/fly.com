package chess

type Point struct {
	Row, Col int
}

func NewPoint(Row, Col int) *Point {
	ret := new(Point)
	ret.Row = Row
	ret.Col = Col
	return ret
}
