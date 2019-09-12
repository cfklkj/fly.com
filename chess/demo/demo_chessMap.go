package main

import "../../chess"

func main() {
	chessMap := chess.NewChessMap()
	chessMap.Board.ShowChessBoarPt()
	chessMap.LoadEndgame("../map/map1.md")
	chessMap.Board.ShowChessBoar()
	//pieces := chess.NewChessPieces()
}
