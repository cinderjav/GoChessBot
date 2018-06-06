package main

type PieceMove struct {
	move  Move
	piece IChessPiece
	score int
}

var kingKills = 0

func minMax(moveMapping map[IChessPiece][]Move, board [8][8]string, turn string, level int, initialTurn string, parentMove Move, parentPiece IChessPiece) (int, Move, IChessPiece) {
	if level == 0 {
		score := analyzeBoard(board, turn, initialTurn)
		return score, parentMove, parentPiece
	}

	// if isKingKilled(board, turn) {
	// 	kingKills += 1
	// 	if turn == initialTurn {
	// 		return -10000
	// 	} else {
	// 		return 10000
	// 	}
	// }

	prunedMap := pruneMinMax(moveMapping, board, turn)
	if turn == initialTurn {
		var bestResult = -1000000000
		var maxMove Move
		var maxPiece IChessPiece
		//var bestMoveChoice Move
		for piece, moves := range prunedMap {
			for _, move := range moves {
				resultMax, _, _ := makeMoveAndGenerate(move, piece, turn, board, level-1, initialTurn)
				bestResult, maxMove, maxPiece = max(bestResult, resultMax, maxMove, move, maxPiece, piece)
			}
		}
		return bestResult, maxMove, maxPiece
	} else {
		var worstResult = 1000000000
		var minMove Move
		var minPiece IChessPiece
		for piece, moves := range prunedMap {
			for _, move := range moves {
				resultMin, _, _ := makeMoveAndGenerate(move, piece, turn, board, level-1, initialTurn)
				worstResult, minMove, minPiece = min(worstResult, resultMin, minMove, move, minPiece, piece)
			}
		}
		return worstResult, minMove, minPiece
	}
}

func max(currentmax, newValue int, currentMove, move Move, currentPiece, piece IChessPiece) (int, Move, IChessPiece) {
	if newValue > currentmax {
		return newValue, move, piece
	}
	return currentmax, currentMove, currentPiece
}

func min(currentmin, newValue int, currentMove, move Move, currentPiece, piece IChessPiece) (int, Move, IChessPiece) {
	if newValue < currentmin {
		return newValue, move, piece
	}
	return currentmin, currentMove, currentPiece
}

var count = 0

func analyzeBoard(board [8][8]string, turn string, initialTurn string) int {
	count += 1
	//always analyze current turn here
	analysisGame := ChessGame{board, turn}
	whitePieces := analysisGame.getWhitePieces()
	blackPieces := analysisGame.getBlackPieces()
	whiteValue := analysisGame.getBoardValueForPieces(whitePieces)
	blackValue := analysisGame.getBoardValueForPieces(blackPieces)
	if initialTurn == WhiteTurn {
		return whiteValue - blackValue
	} else {
		return blackValue - whiteValue
	}
}

func makeMoveAndGenerate(move Move, piece IChessPiece, turn string, board [8][8]string, level int, initialTurn string) (int, Move, IChessPiece) {
	newBoard := makeBoardMove(piece, move, board)
	newTurn := getNextPlayerTurn(turn)
	newChessGame := ChessGame{newBoard, newTurn}
	pieces := newChessGame.getPiecesForTurn()
	newMapping := getAllAvailableMovesForTurn(pieces, &newChessGame)
	return minMax(newMapping, newBoard, newTurn, level, initialTurn, move, piece)
}

func pruneMinMax(moveMapping map[IChessPiece][]Move, board [8][8]string, turn string) map[IChessPiece][]Move {
	prunedMap := make(map[IChessPiece][]Move)
	var pruneList []Move
	for piece, moves := range moveMapping {
		for _, move := range moves {
			prune, pruneAll := pruneMoveMinMax(piece, move, board, turn)
			if pruneAll {
				prunedMapAll := make(map[IChessPiece][]Move)
				prunedMapAll[piece] = append(prunedMapAll[piece], move)
				return prunedMapAll
			}
			if prune {
				pruneList = append(pruneList, move)
			} else {
				prunedMap[piece] = append(prunedMap[piece], move)
			}
		}
	}

	return prunedMap
	// for _, badMove := range pruneMoves {
	// 	delete(moveMapping, badPiece)
	// }

}

func pruneMoveMinMax(piece IChessPiece, move Move, board [8][8]string, turn string) (bool, bool) {
	//change defending moves to simulate the move first
	postMoveBoard := makeBoardMove(piece, move, board)
	defending := isEnemyDefendingMove(move, turn, postMoveBoard)
	//Remove any move that will cause King to die next turn
	//should be handled in does my move lead to checkmate
	pruneKingChecked := pruneMoveKingChecked(piece, turn, board, move, defending)
	if pruneKingChecked {
		return true, false
	}
	//dont kill a piece that is defended and lower value than you

	if isCheckmateMove(piece, move, turn, board, defending, -1) {

		//prune all other moves in map, this is clearly the best move
		//idea is to limit score for this path, should result in much lower score propagated up to original move
		return false, true
	}

	//does my move lead to a checkmate?
	newb := makeBoardMove(piece, move, board)
	next := getNextPlayerTurn(turn)
	oppChess := ChessGame{newb, next}
	oppopieces := oppChess.getPiecesForTurn()
	newMovesMapping := getAllAvailableMovesForTurn(oppopieces, &oppChess)
	for piecekey, movesvalue := range newMovesMapping {
		for _, moveValue := range movesvalue {
			postMoveInner := makeBoardMove(piecekey, moveValue, newb)
			if isKingKilled(postMoveInner, turn) {
				return true, false
			}
			//is king on board check for black here
			//keep condition below also
			defendingInner := isEnemyDefendingMove(moveValue, next, postMoveInner)
			if isCheckmateMove(piecekey, moveValue, next, newb, defendingInner, -1) {
				return true, false
			}
		}
	}

	return false, false
}
