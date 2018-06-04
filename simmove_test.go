package main

import "testing"

func TestKingChecked(t *testing.T) {
	fen := "1N1kB3/1pr5/3p1p2/P3N3/4R3/2p5/K6q/1n6 w - - 0 1"
	fenObject := FenRequest{fen}
	v := Run(fenObject)
	println(v)
	if v != "a2-b3" {
		t.Error("Expected h2-b2, got ", v)
	}
}

func TestFindCheckmate(t *testing.T) {
	fen := "1N1kB3/1pr5/3p1p2/P3N3/4R3/2p5/7q/1K6 b - - 0 1"
	fenObject := FenRequest{fen}
	v := Run(fenObject)
	println(v)
	if v != "h2-b2" {
		t.Error("Expected h2-b2, got ", v)
	}
}

func TestToughSpot(t *testing.T) {
	fen := "1N1kr3/1pr5/3p1pB1/P3N3/4R3/2p5/K6P/1n4q1 w - - 0 1"
	fenObject := FenRequest{fen}
	v := Run(fenObject)
	println(v)
	if v != "g6-e8" {
		t.Error("Expected g6-e8, got ", v)
	}
}

//1N1kB3/1pr5/3p1p2/P3N3/4R3/2p5/7q/1K6 b - - 0 1
