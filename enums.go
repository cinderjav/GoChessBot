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
	PawnScore       = 3
	BishopScore     = 6
	KnightScore     = 6
	RookScore       = 10
	QueenScore      = 20
	KingScore       = 100
)

var MaxRecursiveLevel = 3
