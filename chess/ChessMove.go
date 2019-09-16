package chess

import "fmt"

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

//移动
func (c *ChessMove) Move(chess ChessStatu, pt Point) *RelationPoints {

	direct, offset := c.getDirection(chess, pt)
	offset *= -1
	pieces := chess.GetPieces()
	fmt.Println(pieces.ChessString(), direct, offset)
	switch pieces.ChessString() {
	case "军":
		return c._Ju(chess, direct, offset)
	case "将":
		return c._Jiang(chess, direct)
	case "炮":
		return c._Pao(chess, direct, offset)
	case "兵":
		return c._Bing(chess, direct)
	case "马":
		return c._Ma(chess, direct)
	case "象":
		return c._Xiang(chess, direct)
	case "仕":
		return c._Bing(chess, direct)
	}
	return nil
}

//获取移动方向
func (c *ChessMove) getDirection(chess ChessStatu, pt Point) (int, int) { //direct offset
	fmt.Println("getDirection", pt, chess.pt)
	if pt.Row == 0 {
		if chess.pt.Col > pt.Col {
			return L, chess.pt.Col - pt.Col
		} else {
			return R, chess.pt.Col - pt.Col
		}
	}
	if pt.Col == 0 {
		if chess.pt.Row > pt.Row {
			return U, chess.pt.Row - pt.Row
		} else {
			return D, chess.pt.Row - pt.Row
		}
	}
	if chess.pt.Row > pt.Row {
		if chess.pt.Col > pt.Col {
			return LU, 0
		} else {
			return LD, 0
		}
	} else {
		if chess.pt.Col > pt.Col {
			return RU, 0
		} else {
			return RD, 0
		}
	}
}

//军
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
