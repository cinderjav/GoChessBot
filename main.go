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
	mux.Handle("/movev3", http.HandlerFunc(handlePlayv3))
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

func handlePlayv3(w http.ResponseWriter, req *http.Request) {
	//"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

	var fenObject FenRequest
	fmt.Println(req.Body)
	err := json.NewDecoder(req.Body).Decode(&fenObject)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println(fenObject)
	move := RunV3(fenObject)
	w.Write([]byte(move))
	// fmt.Println(move)
	// moveJson, err := json.Marshal(move)
	// if err != nil {
	// 	panic(err)
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(moveJson)
}

func Run(fenObject FenRequest) string {
	board, turn := fenParser(fenObject.Fen)
	chessGame := getChessGame(board, turn)
	executedMove := chessGame.executeMove()
	return executedMove
}

func RunV3(fenObject FenRequest) string {
	board, turn := fenParser(fenObject.Fen)
	chessGame := getChessGame(board, turn)
	executedMove := chessGame.executeMoveMinMax()
	return executedMove
}

type FenRequest struct {
	Fen string
}
