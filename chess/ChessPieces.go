package chess

import "fmt"

type PiecesName int

type ChessStatu struct {
	pieces PiecesName
	pt     Point
	deaded bool
}

type ChessPieces struct {
	pieces map[PiecesName]ChessStatu
}

//获取棋子名称
func (c *ChessStatu) GetPieces() PiecesName {
	return c.pieces
}

//新建棋子类
func NewChessPieces() *ChessPieces {
	ret := new(ChessPieces)
	ret.pieces = make(map[PiecesName]ChessStatu)
	ret.buildChessPieces()
	return ret
}

//移动棋子
func (c *ChessStatu) GetOffsetPoint(offsetLevel int, offsetVertical int) Point {
	var pt Point
	pt.X = c.pt.X + offsetLevel
	pt.Y = c.pt.Y + offsetVertical
	return pt
}

//是否是己方棋子
func (c *ChessStatu) IsOwnPieces(pieces PiecesName) bool {
	if pieces == Null { //没有棋子
		return false
	}
	if c.pieces <= R_jiang && pieces <= R_jiang { //红方
		return true
	}
	if c.pieces > R_jiang && pieces > R_jiang { //黑方
		return true
	}
	return false //敌对方
}

//新建棋子
func (c *ChessPieces) buildChessPieces() {
	for i := R_ju1; i <= B_bin5; i++ {
		var chessStatu ChessStatu
		chessStatu.pieces = PiecesName(i)
		c.pieces[chessStatu.pieces] = chessStatu
	}
	fmt.Println("BuildChessPieces", c.pieces)
}
