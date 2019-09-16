package chess

import "fmt"

type ChessStatu struct {
	pieces PiecesName
	pt     Point
	deaded bool //已亡
	used   bool //已使用
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
	pt.Row = c.pt.Row + offsetVertical //垂直  行
	pt.Col = c.pt.Col + offsetLevel    //水平  列
	return pt
}

//是否是己方棋子
func (c *ChessStatu) IsOwnPieces(pieces PiecesName) bool {
	if pieces == Null { //没有棋子
		return false
	}
	if c.pieces <= R_bin5 && pieces <= R_bin5 { //红方
		return true
	}
	if c.pieces > R_bin5 && pieces > R_bin5 { //黑方
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

//摆设棋子
func (c *ChessPieces) SetPoint(pieces PiecesName, pt Point) bool {
	chessStatu, ok := c.pieces[pieces]
	if ok {
		chessStatu.pt = pt
		chessStatu.pieces = pieces
		chessStatu.deaded = false
		return true
	}
	return false
}

//吃掉棋子
func (c *ChessPieces) EatUp(pieces PiecesName) bool {
	chessStatu, ok := c.pieces[pieces]
	if ok {
		chessStatu.deaded = true
		return true
	}
	return false
}

//获取棋子信息
func (c *ChessPieces) GetPiecesStatu(pieces PiecesName) *ChessStatu {
	chessStatu, ok := c.pieces[pieces]
	if ok {
		return &chessStatu
	}
	return nil

}

//获取可用棋子
func (c *ChessPieces) GetPieces(piecesName string, isRed bool) PiecesName {
	for pieces, statu := range c.pieces {
		if pieces.IsThisPieces(piecesName) && (isRed == pieces.IsRed() || !isRed == pieces.IsBlack()) {
			if !statu.deaded {
				return pieces
			}
		}
	}
	return Null
}
