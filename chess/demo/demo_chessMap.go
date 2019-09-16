package main

import "../../chess"

func main() {
	chessMap := chess.NewChessMap()
	chessMap.Board.ShowChessBoarPt()
	chessMap.LoadEndgame("../map/map1.md")
	chessMap.Board.ShowChessBoar()
	chessMap.MoveChess(chess.Point{0, 0}, chess.Point{3, 0})
	chessMap.Board.ShowChessBoar()
	//chessMap.MoveChess(chess.Point{0, 0}, chess.Point{3, 5})
	//chessMap.Board.ShowChessBoar()
	//pieces := chess.NewChessPieces()
}
