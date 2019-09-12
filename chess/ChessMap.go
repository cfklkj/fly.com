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
}

func NewChessMap() *ChessMap {
	ret := new(ChessMap)
	ret.Board = NewChessBoard()
	ret.Pieces = NewChessPieces()
	return ret
}

//将棋子设置到棋盘
func (c *ChessMap) setToBoar(pieces PiecesName, pt Point) bool {
	if !c.Board.FindPoint(pt) {
		return false
	}
	c.Board.SetPieces(pieces, pt)
	c.Pieces.SetPoint(pieces, pt)
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
