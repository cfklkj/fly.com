package chess

//piecesName
var g_pieceName = []string{"", "军", "马", "象", "仕", "将", "仕", "象", "马", "军", "炮", "炮", "兵", "兵", "兵", "兵", "兵",
	"军", "马", "象", "仕", "将", "仕", "象", "马", "军", "炮", "炮", "兵", "兵", "兵", "兵", "兵"}

const (
	OverStep = -1 //越界
	Null     = 0  //无棋子
	//红棋
	R_ju1    = 1
	R_ma1    = 2
	R_xiang1 = 3
	R_shi1   = 4
	R_jiang  = 5
	R_shi2   = 6
	R_xiang2 = 7
	R_ma2    = 8
	R_ju2    = 9
	R_pao1   = 10
	R_pao2   = 11
	R_bin1   = 12
	R_bin2   = 13
	R_bin3   = 14
	R_bin4   = 15
	R_bin5   = 16
	//黑棋
	B_ju1    = 17
	B_ma1    = 18
	B_xiang1 = 19
	B_shi1   = 20
	B_jiang  = 21
	B_shi2   = 22
	B_xiang2 = 23
	B_ma2    = 24
	B_ju2    = 25
	B_pao1   = 26
	B_pao2   = 27
	B_bin1   = 28
	B_bin2   = 29
	B_bin3   = 30
	B_bin4   = 31
	B_bin5   = 32
)

const (
	L  = 1
	R  = 2
	U  = 3
	D  = 4
	LU = 5
	LD = 6
	RU = 7
	RD = 8
)

type PiecesName int

//得到象棋名称
func (c *PiecesName) ChessString() string {
	ChessName := *c
	if B_bin5 < ChessName || ChessName < R_ju1 {
		return ""
	}
	return g_pieceName[ChessName]
}

//判断象棋名称
func (c *PiecesName) IsThisPieces(piecesName string) bool {
	ChessName := *c
	if B_bin5 < ChessName || ChessName < R_ju1 {
		return false
	}
	return g_pieceName[ChessName] == piecesName
}

//判断红黑
func (c *PiecesName) IsRed() bool {
	return !c.IsBlack()
}

func (c *PiecesName) IsBlack() bool {
	ChessName := *c
	if B_bin5 < ChessName || ChessName < R_ju1 {
		return false
	}
	return ChessName > R_bin5
}
