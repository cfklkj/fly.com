package chess

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
	return ret
}

//设置点
func (c *BoardPoint) addPoint(x, y int) {
	point := NewPoint(x, y)
	c.pt[*point] = PiecesName(x + y)
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
func (c *BoardPoint) GetPointPieces(pt Point) PiecesName {
	value, ok := c.pt[pt]
	if ok {
		return value
	}
	return OverStep //越界
}

//将-边界
func (c *ChessBoard) board_jing() {
	for i := 0; i <= 2; i++ {
		for j := 3; j <= 5; j++ {
			c.points_B_jiang.addPoint(i, j)
			c.points_R_jiang.addPoint(i+7, j+7)
		}
	}
}
func (c *ChessBoard) FindPoint_jiang(pieces PiecesName, pt Point) bool {
	if pieces > R_jiang { //黑
		return c.points_B_jiang.findPoint(pt)
	} else { //红
		return c.points_R_jiang.findPoint(pt)
	}
}
func (c *ChessBoard) FindPoint_bian(chess *ChessStatu, pt Point) bool {
	if chess.pieces > R_jiang { //黑
		return chess.pt.Y >= pt.Y
	} else { //红
		return chess.pt.Y <= pt.Y
	}
}

//棋盘
func (c *ChessBoard) board_chessBoard() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 9; j++ {
			c.points_chessBoard.addPoint(i, j)
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
func (c *ChessBoard) FindPoint_shi(pieces PiecesName, pt Point) bool {
	if pieces > R_jiang { //黑
		return c.points_B_shi.findPoint(pt)
	} else { //红
		return c.points_R_shi.findPoint(pt)
	}
}
func (c *ChessBoard) FindPoint_xiang(pieces PiecesName, pt Point) bool {
	if pieces > R_jiang { //黑
		return c.points_B_xiang.findPoint(pt)
	} else { //红
		return c.points_R_xiang.findPoint(pt)
	}
}

//创建棋盘
func (c *ChessBoard) BuildBoard() {
	c.board_chessBoard()
	c.board_jing()
	c.limitPoints_shi()
	c.limitPoints_xiang()
}

//查找棋盘点
func (c *ChessBoard) findPoint(point Point) bool {
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
