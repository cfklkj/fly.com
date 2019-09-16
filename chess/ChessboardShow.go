package chess

import (
	"container/list"
	"fmt"
)

var g_listChessBoard *list.List

func InstanceListBoard() *list.List {
	if g_listChessBoard == nil {
		g_listChessBoard = list.New()
	}
	return g_listChessBoard
}

func (c *BoardPoint) AddTolist(pt *Point) {
	InstanceListBoard().PushBack(pt)
}

//显示象棋
func (c *ChessBoard) showBoar(pts *BoardPoint) {
	i := 1
	rows := 0
	for pt := InstanceListBoard().Front(); pt != nil; pt = pt.Next() { //依据链表进行
		var Pt *Point = pt.Value.(*Point) //类型强转
		pieces := pts.GetPointPieces(*Pt)
		if i > 1 {
			fmt.Print(" - ")
		}
		//fmt.Printf("%2d", *Pt)
		if pieces == 0 {
			fmt.Printf("  ")
		} else {
			fmt.Printf(pieces.ChessString())
		}
		i++
		if i == 10 {
			fmt.Println("")
			rows++
			if rows == 5 {
				i = 1
				fmt.Print("---------------楚河    汉界--------------")
				fmt.Println("")
				continue
			}
			if rows == 10 {
				continue
			}
			for ; i > 1; i-- {
				fmt.Print("|    ")
			}
			fmt.Println("")
		}
	}
}

//显示点
func (c *ChessBoard) showChessPt(pts *BoardPoint) {
	i := 1
	rows := 0
	for pt := InstanceListBoard().Front(); pt != nil; pt = pt.Next() { //依据链表进行
		var Pt Point = *pt.Value.(*Point) //类型强转
		if i > 1 {
			fmt.Print(" - ")
		}
		Pt.Row += 1
		Pt.Col += 1
		fmt.Printf("%1d", Pt)

		i++
		if i == 10 {
			fmt.Println("")
			rows++
			if rows == 5 {
				i = 1
				fmt.Print("---------------              楚河    汉界              --------------")
				fmt.Println("")
				continue
			}
			if rows == 10 {
				continue
			}
			for ; i > 1; i-- {
				fmt.Print("  |     ")
			}
			fmt.Println("")
		}
	}
}

func (c *ChessBoard) ShowChessBoar() {
	c.showBoar(&c.points_chessBoard)
}
func (c *ChessBoard) ShowChessBoarPt() {
	c.showChessPt(&c.points_chessBoard)
}
