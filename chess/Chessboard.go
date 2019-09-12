package chess

import "fmt"

//棋盘点结构
type BoardPoint struct {
	pt map[Point]PiecesName
}

//棋盘布局
type ChessBoard struct {
	points_chessBoard BoardPoint
	points_B_jiang    BoardPoint
	points_R_jiang    BoardPoint
	points_B_shi      BoardPoint
	points_R_shi      BoardPoint
	points_B_xiang    BoardPoint
	points_R_xiang    BoardPoint
}

func NewChessBoard() *ChessBoard {
	ret := new(ChessBoard)
	ret.points_chessBoard.pt = make(map[Point]PiecesName)
	ret.points_B_jiang.pt = make(map[Point]PiecesName)
	ret.points_R_jiang.pt = make(map[Point]PiecesName)
	ret.points_B_shi.pt = make(map[Point]PiecesName)
	ret.points_R_shi.pt = make(map[Point]PiecesName)
	ret.points_B_xiang.pt = make(map[Point]PiecesName)
	ret.points_R_xiang.pt = make(map[Point]PiecesName)
	ret.buildBoard()
	return ret
}

//设置点
func (c *BoardPoint) addPoint(x, y int) {
	point := NewPoint(x, y)
	c.pt[*point] = PiecesName(x + y)
}
func (c *BoardPoint) addPointTolist(x, y int) {
	point := NewPoint(x, y)
	c.pt[*point] = 0 //PiecesName(x + y)
	c.AddTolist(point)
}
func (c *BoardPoint) addPoints(pts []Point) {
	for _, pt := range pts {
		c.pt[pt] = PiecesName(pt.X + pt.Y)
	}
}

//查找点
func (c *BoardPoint) findPoint(pt Point) bool {
	_, ok := c.pt[pt]
	if ok {
		return true
	}
	return false
}

//设置点棋子
func (c *BoardPoint) SetPointPieces(pt Point, pieces PiecesName) {
	c.pt[pt] = pieces
}

//获取点棋子
func (c *BoardPoint) GetPointPieces(pt Point) PiecesName {
	value, ok := c.pt[pt]
	if ok {
		return value
	}
	return OverStep //越界
}

//将-边界
func (c *ChessBoard) board_jing() {
	for i := 0; i <= 2; i++ { //cow
		for j := 3; j <= 5; j++ { //row
			c.points_B_jiang.addPoint(i, j)
			c.points_R_jiang.addPoint(i+7, j)
		}
	}
}

//将
func (c *ChessBoard) FindPoint_jiang(pieces PiecesName, pt Point) bool {
	if pieces.IsBlack() { //黑
		return c.points_B_jiang.findPoint(pt)
	} else { //红
		return c.points_R_jiang.findPoint(pt)
	}
}

//兵
func (c *ChessBoard) FindPoint_bian(chess *ChessStatu, pt Point) bool {
	if chess.pieces > R_bin5 { //黑
		return chess.pt.Y >= pt.Y
	} else { //红
		return chess.pt.Y <= pt.Y
	}
}

//棋盘
func (c *ChessBoard) board_chessBoard() {
	for i := 0; i < 10; i++ { //row 10
		for j := 0; j < 9; j++ { //cow 8
			c.points_chessBoard.addPointTolist(i, j)
		}
	}
}

//限制点
func (c *ChessBoard) limitPoints_shi() {
	B_pts := []Point{Point{0, 3}, Point{0, 5}, Point{1, 4}, Point{2, 3}, Point{2, 5}}
	c.points_B_shi.addPoints(B_pts)
	R_pts := []Point{Point{9, 3}, Point{9, 5}, Point{8, 4}, Point{7, 3}, Point{7, 5}}
	c.points_R_shi.addPoints(R_pts)
}
func (c *ChessBoard) limitPoints_xiang() {
	B_pts := []Point{Point{0, 2}, Point{0, 6}, Point{2, 0}, Point{2, 8}, Point{5, 2}, Point{5, 6}}
	c.points_B_xiang.addPoints(B_pts)
	R_pts := []Point{Point{9, 2}, Point{9, 6}, Point{7, 0}, Point{7, 8}, Point{6, 2}, Point{6, 6}}
	c.points_R_xiang.addPoints(R_pts)
}

//仕
func (c *ChessBoard) FindPoint_shi(pieces PiecesName, pt Point) bool {
	if pieces.IsBlack() { //黑
		return c.points_B_shi.findPoint(pt)
	} else { //红
		return c.points_R_shi.findPoint(pt)
	}
}

//象
func (c *ChessBoard) FindPoint_xiang(pieces PiecesName, pt Point) bool {
	fmt.Println("sd", pieces, pt, c.points_B_xiang)
	if pieces.IsBlack() { //黑
		return c.points_B_xiang.findPoint(pt)
	} else { //红
		return c.points_R_xiang.findPoint(pt)
	}
}

//创建棋盘
func (c *ChessBoard) buildBoard() {
	c.board_chessBoard()
	c.board_jing()
	c.limitPoints_shi()
	c.limitPoints_xiang()
}

//查找棋盘点
func (c *ChessBoard) FindPoint(point Point) bool {
	return c.points_chessBoard.findPoint(point)
}

//设置棋盘棋子
func (c *ChessBoard) SetPieces(pieces PiecesName, point Point) {
	c.points_chessBoard.SetPointPieces(point, pieces)
}

//获取棋盘点棋子
func (c *ChessBoard) GetPointPieces(point Point) PiecesName {
	return c.points_chessBoard.GetPointPieces(point)
}
