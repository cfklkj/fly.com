package chess

/*
移动棋子关系点
*/
type RelationPoints struct {
	directPt Point
	checkPts []Point
}
type ChessMove struct {
	pts RelationPoints
}

func NewChessMove() *ChessMove {
	ret := new(ChessMove)
	return ret
}

//马
func (c *ChessMove) _Ju(chess ChessStatu, direction int, offset int) *RelationPoints {
	c.pts.checkPts = []Point{}
	switch direction {
	case L:
	case R:
		c.pts.directPt = chess.GetOffsetPoint(offset, 0)
		for i := offset - 1; i > 0; i-- { //路径点
			point := chess.GetOffsetPoint(i, 0)
			c.pts.checkPts = append(c.pts.checkPts, point)
		}
	case U:
	case D:
		c.pts.directPt = chess.GetOffsetPoint(0, offset)
		for i := offset - 1; i > 0; i-- { //路径点
			point := chess.GetOffsetPoint(0, i)
			c.pts.checkPts = append(c.pts.checkPts, point)
		}
	default:
		return nil
	}
	return &c.pts
}

//马
func (c *ChessMove) _Ma(chess ChessStatu, direction int) *RelationPoints {
	switch direction {
	case LU:
		c.pts.directPt = chess.GetOffsetPoint(-1, 2)
		point := chess.GetOffsetPoint(0, 1) //马脚
		c.pts.checkPts = []Point{point}
	case LD:
		c.pts.directPt = chess.GetOffsetPoint(-1, -2)
		point := chess.GetOffsetPoint(0, -1) //马脚
		c.pts.checkPts = []Point{point}
	case RU:
		c.pts.directPt = chess.GetOffsetPoint(1, 2)
		point := chess.GetOffsetPoint(0, 1) //马脚
		c.pts.checkPts = []Point{point}
	case RD:
		c.pts.directPt = chess.GetOffsetPoint(1, -2)
		point := chess.GetOffsetPoint(0, -1) //马脚
		c.pts.checkPts = []Point{point}
	default:
		return nil
	}
	return &c.pts
}

//象
func (c *ChessMove) _Xiang(chess ChessStatu, direction int) *RelationPoints {
	switch direction {
	case LU:
		c.pts.directPt = chess.GetOffsetPoint(-2, 2)
		point := chess.GetOffsetPoint(-1, 1) //象脚
		c.pts.checkPts = []Point{point}
	case LD:
		c.pts.directPt = chess.GetOffsetPoint(-2, -2)
		point := chess.GetOffsetPoint(-1, -1)
		c.pts.checkPts = []Point{point}
	case RU:
		c.pts.directPt = chess.GetOffsetPoint(2, 2)
		point := chess.GetOffsetPoint(1, 1)
		c.pts.checkPts = []Point{point}
	case RD:
		c.pts.directPt = chess.GetOffsetPoint(-2, -2)
		point := chess.GetOffsetPoint(1, -1)
		c.pts.checkPts = []Point{point}
	default:
		return nil
	}
	return &c.pts
}

//仕
func (c *ChessMove) _Shi(chess ChessStatu, direction int) *RelationPoints {
	c.pts.checkPts = []Point{}
	switch direction {
	case LU:
		c.pts.directPt = chess.GetOffsetPoint(-1, 1)
	case LD:
		c.pts.directPt = chess.GetOffsetPoint(-1, -1)
	case RU:
		c.pts.directPt = chess.GetOffsetPoint(1, 1)
	case RD:
		c.pts.directPt = chess.GetOffsetPoint(1, -1)
	default:
		return nil
	}
	return &c.pts
}

//将
func (c *ChessMove) _Jiang(chess ChessStatu, direction int) *RelationPoints {
	return c._Ju(chess, direction, 1)
}

//炮
func (c *ChessMove) _Pao(chess ChessStatu, direction int, offset int) *RelationPoints {
	return c._Ju(chess, direction, offset)
}

//兵
func (c *ChessMove) _Bing(chess ChessStatu, direction int) *RelationPoints {
	return c._Ju(chess, direction, 1)
}
