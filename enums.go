package main

const (
	WhitePawn   = "P"
	WhiteRook   = "R"
	WhiteKnight = "N"
	WhiteBishop = "B"
	WhiteQueen  = "Q"
	WhiteKing   = "K"
	BlackPawn   = "p"
	BlackRook   = "r"
	BlackKnight = "n"
	BlackBishop = "b"
	BlackQueen  = "q"
	BlackKing   = "k"
	EmptySpace  = ""
)

const (
	WhiteTurn = "w"
	BlackTurn = "b"
)

const (
	EmptySpaceScore = 0
	PawnScore       = 1
	BishopScore     = 3
	KnightScore     = 3
	RookScore       = 5
	QueenScore      = 9
	KingScore       = 100
)

var MaxRecursiveLevel = 4
