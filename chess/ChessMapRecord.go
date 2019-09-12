package chess

/*
棋谱设定
*/

type ChessMapRecord struct {
	livedPieces map[PiecesName]int
}

func NewChessMapRecord() *ChessMapRecord {
	ret := new(ChessMapRecord)
	ret.livedPieces = make(map[PiecesName]int)
	return ret
}

func (c *ChessMapRecord) AddPieces(pieces PiecesName) {
	c.livedPieces[pieces] = 1
}
func (c *ChessMapRecord) DelPieces(pieces PiecesName) {
	delete(c.livedPieces, pieces)
}
