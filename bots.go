package main

import (
	"math/rand"
	"time"
)

// RandomBot will make a random move in an empty tile on the
// board it is forced to, based on the previously made move.
func RandomBot(previousMove *Move, board *UltimateBoard) *Move {
	x := previousMove.TileX
	y := previousMove.TileY

	if (x == -1 && y == -1) || board[previousMove.TileX][previousMove.TileY].HasWinner() != EMPTY || len(board[previousMove.TileX][previousMove.TileY].ValidMoves(previousMove.TileX, previousMove.TileY)) == 0 {
		// If:
		//  - This player is starting the game
		//  - The board the bot is supposed to play on is already won
		//  - The board the bot is supposed to play on is already full
		//      -> pick a random tile on a random board instead.
		validMoves := board.ValidMoves()
		randomMoveIndex := rand.Intn(len(validMoves))
		return validMoves[randomMoveIndex]
	}

	// Else, we need to play on the board corresponding to the last move's Tile X/Y position.
	validMoves := board[previousMove.TileX][previousMove.TileY].ValidMoves(previousMove.TileX, previousMove.TileY)
	randomMoveIndex := rand.Intn(len(validMoves))
	return validMoves[randomMoveIndex]
}

// MonteCarloBot uses a Monte Carlo Search Tree to look for the best possible move.
// The function will use TIME_TO_THINK (globally defined) seconds to try to decide
// the "best" next move it can make. A win counts as +1, a tie as 0 and loses as -1,
// each divided with the number of moves it took to reach that state. This means moves
// where few following moves lead to a win are strongly favored, while moves that within
// a few moves will lead to a loss are strongly disfavored.
func MonteCarloBot(playerNumber int, previousMove *Move, board *UltimateBoard) *Move {
	start := time.Now()
	var movesToTry []*Move

	x := previousMove.TileX
	y := previousMove.TileY

	if (x == -1 && y == -1) || board[x][y].HasWinner() != EMPTY || len(board[x][y].ValidMoves(x, y)) == 0 {
		// If:
		//  - The player is making the first move
		//  - The board the bot is supposed to play on is already won
		//  - The board is already full
		//      -> The bot can play on any board.
		movesToTry = board.ValidMoves()

		if len(movesToTry) == 0 {
			// HackerRank does not properly detect when a game is already
			// tied, but will force players to fill up all the boards
			// before calling it, so we keep playing...
			movesToTry = board.AllPossibleMoves()
		}
	} else {
		// Otherwise we settle for the moves we can
		// make on the board the bot was "forced to".
		movesToTry = board[x][y].ValidMoves(x, y)
	}

	ties := make(map[Move]float64)
	wins := make(map[Move]float64)
	weightedWins := make(map[Move]float64)
	losses := make(map[Move]float64)
	weightedLosses := make(map[Move]float64)
	gamesPlayed := 0

	allValidMoves := make([]Move, len(movesToTry))
	for i := 0; i < len(movesToTry); i++ {
		allValidMoves[i] = *movesToTry[i]
	}

	for {
		// Until we run out of time...
		for _, move := range movesToTry {
			gamesPlayed += 1

			// Keep track of whose turn it is
			var previousPlayer int
			// Keep track of how many moves were needed to end the game
			var movesUntilGameEnded float64

			// Create a copy of the board
			localBoard := board.Copy()
			localMove := move.Copy()

			// It's the bot's turn, so the previous player must have made
			// the last move -> Make the move on the board.
			if playerNumber == 1 {
				localBoard[localMove.BoardX][localMove.BoardY][localMove.TileX][localMove.TileY] = PLAYER_1_CONTROLLED
				previousPlayer = 1
			} else {
				localBoard[localMove.BoardX][localMove.BoardY][localMove.TileX][localMove.TileY] = PLAYER_2_CONTROLLED
				previousPlayer = 2
			}
			movesUntilGameEnded += 1.0

			// Simulate the rest of the game with two RandomBots.
			for localBoard.HasWinner() == EMPTY {
				if len(localBoard.ValidMoves()) == 0 {
					ties[*move] += 1.0
					break
				}

				movesUntilGameEnded += 1.0

				localMove = RandomBot(localMove, localBoard)
				if previousPlayer == 1 {
					localBoard[localMove.BoardX][localMove.BoardY][localMove.TileX][localMove.TileY] = PLAYER_2_CONTROLLED
					previousPlayer = 2
				} else if previousPlayer == 2 {
					localBoard[localMove.BoardX][localMove.BoardY][localMove.TileX][localMove.TileY] = PLAYER_1_CONTROLLED
					previousPlayer = 1
				}
			}

			localBoardWinner := localBoard.HasWinner()
			if (playerNumber == 1 && localBoardWinner == PLAYER_1_CONTROLLED) || (playerNumber == 2 && localBoardWinner == PLAYER_2_CONTROLLED) {
				wins[*move] += 1.0
				weightedWins[*move] += (1.0 / movesUntilGameEnded)
			} else if (playerNumber == 1 && localBoardWinner == PLAYER_2_CONTROLLED) || (playerNumber == 2 && localBoardWinner == PLAYER_1_CONTROLLED) {
				losses[*move] += 1.0
				weightedLosses[*move] += (1.0 / movesUntilGameEnded)
			}
		}

		// Break when the bot runs out of time
		end := time.Now()
		if end.Sub(start).Seconds() > TIME_TO_THINK {
			//fmt.Printf("MonteCarloBot had time to play %d simulated games using %d valid moves (~%d per valid move) before running out of time!\n", gamesPlayed, len(movesToTry), (gamesPlayed / len(movesToTry)))
			break
		}
	}

	bestScore := -10000.0
	var bestMove Move

	for _, move := range allValidMoves {
		score := (weightedWins[move] - (weightedLosses[move] * 2.0)) / (wins[move] + losses[move] + ties[move])
		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}

	// fmt.Printf("The best move (%s) had a score of %f\n", bestMoveString, bestScore)
	return &bestMove
}
