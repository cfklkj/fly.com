package chess

import "fmt"

/*
走法限制
*/
type Chess struct {
	directPt Point
	checkPts []Point
}
type ChessRun struct {
	pts RelationPoints
}

func NewChessRun() *ChessRun {
	ret := new(ChessRun)
	return ret
}

//移动判断
func (c *ChessRun) MoveCheck(chess *ChessStatu, board *ChessBoard, points *RelationPoints) bool {
	pieces := chess.GetPieces()
	switch pieces.ChessString() {
	case "军":
		return c._Ju(chess, board, points)
	case "将":
		return c._Jiang(chess, board, points)
	case "炮":
		return c._Pao(chess, board, points)
	case "兵":
		return c._Bing(chess, board, points)
	case "马":
		return c._Ma(chess, board, points)
	case "象":
		return c._Xiang(chess, board, points)
	case "仕":
		return c._Bing(chess, board, points)
	}
	return false
}

//军
func (c *ChessRun) _Ju(chess *ChessStatu, board *ChessBoard, points *RelationPoints) bool {

	directPieces := board.GetPointPieces(points.directPt)
	if directPieces == OverStep { //越界
		fmt.Print("_Ju1")
		return false
	}
	if chess.IsOwnPieces(directPieces) { //目标点
		fmt.Print("_Ju2")
		return false
	}
	//路径点
	for _, pt := range points.checkPts {
		pieces := board.GetPointPieces(pt)
		if pieces != Null {
			fmt.Print("_Ju3")
			return false
		}
	}
	return true
}

//马
func (c *ChessRun) _Ma(chess *ChessStatu, board *ChessBoard, points *RelationPoints) bool {
	return c._Ju(chess, board, points)
}

//象
func (c *ChessRun) _Xiang(chess *ChessStatu, board *ChessBoard, points *RelationPoints) bool {
	directPieces := board.GetPointPieces(points.directPt)
	if directPieces == OverStep { //越界
		return false
	}
	if chess.IsOwnPieces(directPieces) { //目标点
		return false
	}
	if !board.FindPoint_xiang(chess.GetPieces(), points.directPt) { //目标点越界
		return false
	}
	//路径点
	for _, pt := range points.checkPts {
		pieces := board.GetPointPieces(pt)
		if pieces != Null {
			return false
		}
	}
	return true
}

//仕
func (c *ChessRun) _Shi(chess *ChessStatu, board *ChessBoard, points *RelationPoints) bool {
	directPieces := board.GetPointPieces(points.directPt)
	if directPieces == OverStep { //越界
		return false
	}
	if chess.IsOwnPieces(directPieces) { //目标点 自己
		return false
	}
	if !board.FindPoint_shi(chess.GetPieces(), points.directPt) { //目标点越界
		return false
	}
	//路径点
	for _, pt := range points.checkPts {
		pieces := board.GetPointPieces(pt)
		if pieces != Null {
			return false
		}
	}
	return true
}

//将
func (c *ChessRun) _Jiang(chess *ChessStatu, board *ChessBoard, points *RelationPoints) bool {
	directPieces := board.GetPointPieces(points.directPt)
	if directPieces == OverStep { //越界
		return false
	}
	if chess.IsOwnPieces(directPieces) { //目标点 自己
		return false
	}
	if !board.FindPoint_jiang(chess.GetPieces(), points.directPt) { //目标点越界
		return false
	}
	//路径点
	for _, pt := range points.checkPts {
		pieces := board.GetPointPieces(pt)
		if pieces != Null {
			return false
		}
	}
	return true
}

//炮
func (c *ChessRun) _Pao(chess *ChessStatu, board *ChessBoard, points *RelationPoints) bool {
	directPieces := board.GetPointPieces(points.directPt)
	if directPieces == OverStep { //越界
		return false
	}
	if chess.IsOwnPieces(directPieces) { //目标点 自己
		return false
	}
	//路径点
	count := 0
	for _, pt := range points.checkPts {
		pieces := board.GetPointPieces(pt)
		if pieces != Null {
			count++
		}
		if count > 1 {
			return false
		}
	}
	if count != 1 { //没有桥
		return false
	}
	return true
}

//兵
func (c *ChessRun) _Bing(chess *ChessStatu, board *ChessBoard, points *RelationPoints) bool {
	//方向错误
	if !board.FindPoint_bian(chess, points.directPt) {
		return false
	}
	directPieces := board.GetPointPieces(points.directPt)
	if directPieces == OverStep { //越界
		return false
	}
	if chess.IsOwnPieces(directPieces) { //目标点 自己
		return false
	}
	return true
}
