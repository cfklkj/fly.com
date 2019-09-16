package chess

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
棋谱设定
*/

type ChessMap struct {
	Board  *ChessBoard
	Pieces *ChessPieces
	Record *ChessMapRecord
	Move   *ChessMove
	Run    *ChessRun
}

func NewChessMap() *ChessMap {
	ret := new(ChessMap)
	ret.Board = NewChessBoard()
	ret.Pieces = NewChessPieces()
	ret.Record = NewChessMapRecord()
	ret.Move = NewChessMove()
	ret.Run = NewChessRun()
	return ret
}

//将棋子设置到棋盘
func (c *ChessMap) setToBoar(pieces PiecesName, pt Point) bool {
	if !c.Board.FindPoint(pt) {
		return false
	}
	c.delFromBoar(pt)
	c.Board.SetPieces(pieces, pt)
	c.Pieces.SetPoint(pieces, pt)
	c.Record.AddPieces(pieces)
	return true
}

//移除棋子
func (c *ChessMap) delFromBoar(pt Point) bool {
	pieces := c.Board.GetPointPieces(pt)
	if pieces == OverStep {
		return false
	}
	fmt.Println("delFromBoar", pieces)
	c.Pieces.EatUp(pieces)
	c.Record.DelPieces(pieces)
	return true
}

//将棋子设置到棋盘  走棋
func (c *ChessMap) MoveChess(srcPt Point, directPt Point) bool {
	if directPt == srcPt {
		fmt.Print("1")
		return false
	}
	if !c.Board.FindPoint(directPt) {
		fmt.Print("1")
		return false
	}
	pieces := c.Board.GetPointPieces(srcPt)
	if pieces == OverStep {
		fmt.Print("1")
		return false
	}
	statu := c.Pieces.GetPiecesStatu(pieces)
	if statu == nil {
		fmt.Print("2")
		return false
	}
	relationPts := c.Move.Move(*statu, directPt) //
	if relationPts == nil {                      //移动位置错误
		fmt.Println("3", *statu, directPt)
		return false
	}
	if relationPts.directPt != directPt { //目标点检测错误
		fmt.Print("4", relationPts.directPt, directPt)
		return false
	}
	if !c.Run.MoveCheck(statu, c.Board, relationPts) { //不可移动到目标点
		fmt.Print("5")
		return false
	}
	c.Board.DelPointPieces(srcPt)
	c.setToBoar(pieces, directPt)
	return true
}

//残局结构体
type LoadEndgameInfo struct {
	Red   []string `json:"红"`
	Black []string `json:"黑"`
}

//加载残局文件
func (c *ChessMap) LoadEndgame(filePath string) bool {

	filePtr, err := os.Open(filePath)
	if err != nil {
		return false
	}
	defer filePtr.Close()

	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	var endGame LoadEndgameInfo
	err = decoder.Decode(&endGame)
	if err != nil {
		fmt.Println("解析endGame内容失败", err)
		return false
	}
	//
	if !c.makeToBoard(endGame.Red, true) {
		return false
	}
	if !c.makeToBoard(endGame.Black, false) {
		return false
	}
	return err == nil
}

//解析红黑方
func (c *ChessMap) makeToBoard(endGame []string, isRed bool) bool {
	//解析 军,1,1
	fmt.Println("解析endGame内容", endGame)
	for _, chess := range endGame {
		a := strings.Split(chess, ",")
		if len(a) != 3 {
			fmt.Println("解析endGame内容失败")
			return false
		}
		piecesName := a[0]
		x, err := strconv.Atoi(a[1])
		if err != nil {
			fmt.Println("err Atoi", a)
			return false
		}
		y, err2 := strconv.Atoi(a[2])
		if err2 != nil {
			fmt.Println("err Atoi", a)
			return false
		}
		//获取可用棋子名称
		pieces := c.Pieces.GetPieces(piecesName, isRed)
		if pieces == Null {
			fmt.Println("err GetPieces")
			return false
		}
		//设置到棋盘
		x -= 1
		y -= 1
		pt := Point{x, y}
		if !c.setToBoar(pieces, pt) {
			fmt.Println("err setToBoar")
			return false
		}
		//判断位置是否正确
		switch piecesName {
		case "象":
			if !c.Board.FindPoint_xiang(pieces, pt) { //目标点越界
				fmt.Println("err FindPoint_xiang", isRed, pieces, pt)
				return false
			}
		case "仕":
			if !c.Board.FindPoint_shi(pieces, pt) { //目标点越界
				fmt.Println("err FindPoint_shi", pt)
				return false
			}
		case "将":
			if !c.Board.FindPoint_jiang(pieces, pt) { //目标点越界
				fmt.Println("err FindPoint_jiang", pt)
				return false
			}
		}
	}
	return true
}
