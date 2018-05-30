package main

import "testing"

func TestInitialFen(t *testing.T) {
	v := fenParser("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	expected := [8][8]string{
		{BlackRook, BlackKnight, BlackBishop, BlackQueen, BlackKing, BlackBishop, BlackKnight, BlackRook},
		{BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn},
		{EmptySpace, EmptySpace, EmptySpace, EmptySpace, EmptySpace, EmptySpace, EmptySpace, EmptySpace},
		{EmptySpace, EmptySpace, EmptySpace, EmptySpace, EmptySpace, EmptySpace, EmptySpace, EmptySpace},
		{EmptySpace, EmptySpace, EmptySpace, EmptySpace, EmptySpace, EmptySpace, EmptySpace, EmptySpace},
		{EmptySpace, EmptySpace, EmptySpace, EmptySpace, EmptySpace, EmptySpace, EmptySpace, EmptySpace},
		{WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn},
		{WhiteRook, WhiteKnight, WhiteBishop, WhiteQueen, WhiteKing, WhiteBishop, WhiteKnight, WhiteRook},
	}
	if v != expected {
		t.Error("Expected false, got ", v)
	}
}

func TestInGameFen(t *testing.T) {
	v := fenParser("rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2")
	expected := [8][8]string{
		{BlackRook, BlackKnight, BlackBishop, BlackQueen, BlackKing, BlackBishop, BlackKnight, BlackRook},
		{BlackPawn, BlackPawn, EmptySpace, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn},
		{EmptySpace, EmptySpace, EmptySpace, EmptySpace, EmptySpace, EmptySpace, EmptySpace, EmptySpace},
		{EmptySpace, EmptySpace, BlackPawn, EmptySpace, EmptySpace, EmptySpace, EmptySpace, EmptySpace},
		{EmptySpace, EmptySpace, EmptySpace, EmptySpace, WhitePawn, EmptySpace, EmptySpace, EmptySpace},
		{EmptySpace, EmptySpace, EmptySpace, EmptySpace, EmptySpace, WhiteKnight, EmptySpace, EmptySpace},
		{WhitePawn, WhitePawn, WhitePawn, WhitePawn, EmptySpace, WhitePawn, WhitePawn, WhitePawn},
		{WhiteRook, WhiteKnight, WhiteBishop, WhiteQueen, WhiteKing, WhiteBishop, EmptySpace, WhiteRook},
	}
	if v != expected {
		t.Error("Expected false, got ", v)
	}
}
