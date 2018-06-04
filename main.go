package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/move", http.HandlerFunc(handlePlay))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

func handlePlay(w http.ResponseWriter, req *http.Request) {
	//"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

	var fenObject FenRequest
	fmt.Println(req.Body)
	err := json.NewDecoder(req.Body).Decode(&fenObject)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println(fenObject)
	executedMove := Run(fenObject)
	w.Write([]byte(executedMove))
}

func Run(fenObject FenRequest) string {
	board, turn := fenParser(fenObject.Fen)
	chessGame := getChessGame(board, turn)
	executedMove := chessGame.executeMove()
	return executedMove
}

type FenRequest struct {
	Fen string
}
