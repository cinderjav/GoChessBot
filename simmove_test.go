package main

import "testing"

// func TestKingChecked(t *testing.T) {
// 	fen := "1N1kB3/1pr5/3p1p2/P3N3/4R3/2p5/K6q/1n6 w - - 0 1"
// 	fenObject := FenRequest{fen}
// 	v := Run(fenObject)
// 	println(v)
// 	if v != "a2-b3" {
// 		t.Error("Expected a2-b3, got ", v)
// 	}
// }

// func TestFindCheckmate(t *testing.T) {
// 	fen := "1N1kB3/1pr5/3p1p2/P3N3/4R3/2p5/7q/1K6 b - - 0 1"
// 	fenObject := FenRequest{fen}
// 	v := Run(fenObject)
// 	println(v)
// 	if v != "h2-b2" {
// 		t.Error("Expected h2-b2, got ", v)
// 	}
// }

// func TestToughSpot(t *testing.T) {
// 	fen := "1N1kr3/1pr5/3p1pB1/P3N3/4R3/2p5/K6P/1n4q1 w - - 0 1"
// 	fenObject := FenRequest{fen}
// 	v := Run(fenObject)
// 	println(v)
// 	if v != "g6-e8" {
// 		t.Error("Expected g6-e8, got ", v)
// 	}
// }

// func TestLongTime(t *testing.T){
// 	//rn1qkbnr/pbpppp1p/6p1/1p1P4/3Q2P1/8/PPP1PP1P/RNB1KBNR b KQkq - 1 4
// 	println("here")
// 	fen := "rn1qkbnr/pbpppp1p/6p1/1p1P4/3Q2P1/8/PPP1PP1P/RNB1KBNR b KQkq - 1 4"
// 	fenObject := FenRequest{fen}
// 	v := Run(fenObject)
// 	println(v)
// 	if v != "g6-f6" {
// 		t.Error("Expected g8-f6, got ", v)
// 	}
// }

//rn1qkb1Q/p1pppp1p/6p1/1p6/6P1/8/PPP1P2P/RNB1K1NR b KQq - 0 9
func TestInvalidMoveFixed(t *testing.T) {
	fen := "rn1qkb1Q/p1pppp1p/6p1/1p6/6P1/8/PPP1P2P/RNB1K1NR b KQq - 0 9"
	fenObject := FenRequest{fen}
	v := Run(fenObject)
	println(v)
	if v == "f8-g7" {
		t.Error("Expected another move, that does not result in checkmate, got ", v)
	}
}
