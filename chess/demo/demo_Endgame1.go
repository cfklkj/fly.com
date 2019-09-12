package main

import "../../chess"

func main() {
	chessMap := chess.NewChessMap()
	var pt chess.Point
	pt.X = 1
	pt.Y = 2
	chessMap.SetPiecesToBoard(chess.R_ma1, pt)
	chessMap.Board.ShowChessBoar()
}
