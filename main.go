package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/movev3", http.HandlerFunc(handlePlayv3))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func handlePlayv3(w http.ResponseWriter, req *http.Request) {
	//"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	setupResponse(&w, req)

	if (*req).Method == "OPTIONS" {
		return
	}
	var fenObject FenRequest
	fmt.Println(req.Body)
	err := json.NewDecoder(req.Body).Decode(&fenObject)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println(fenObject)
	move, score := RunV3(fenObject)

	moveResponse := MoveResponse{move, score}
	movejson, err := json.Marshal(moveResponse)
	//w.Write([]byte(move))
	w.Write(movejson)
	// fmt.Println(move)
	// moveJson, err := json.Marshal(move)
	// if err != nil {
	// 	panic(err)
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(moveJson)
}

func RunV3(fenObject FenRequest) (string, int) {
	board, turn := fenParser(fenObject.Fen)
	chessGame := getChessGame(board, turn)
	executedMove, score := chessGame.executeMoveMinMax()
	return executedMove, score
}

type FenRequest struct {
	Fen string
}

type MoveResponse struct {
	Move  string
	Score int
}
