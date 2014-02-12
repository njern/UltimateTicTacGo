package main

// TictactoeBoard represents a single Tic-Tac-Toe board, with
// a grid of 3x3 integers, which can be (EMPTY
// || PLAYER_1_CONTROLLED || PLAYER_2_CONTROLLED)
type TictactoeBoard [3][3]int

// UltimateBoard represents an ultimate Tic-Tac-Toe board, with
// a grid of 3x3 TicTacToeBoards. Each board can be either
// Undecided (EMPTY), won by player 1 (PLAYER_1_CONTROLLED) or won
// by player 2 (PLAYER_2_CONTROLLED)
type UltimateBoard [3][3]TictactoeBoard

// HasWinner returns 1 if player 1 has won the board or 2
// if player 2 has won the board. If the board has not yet
// been won, it returns 0.
func (board *TictactoeBoard) HasWinner() int {
	for i := 0; i < 3; i++ {
		if board[i][0] != EMPTY && board[i][0] == board[i][1] && board[i][0] == board[i][2] {
			/*
			   |X| | |      | |X| |       | | |X|
			   |X| | |  OR  | |X| |  OR   | | |X|
			   |X| | |      | |X| |       | | |X|
			*/
			return board[i][0]
		}

		if board[0][i] != EMPTY && board[0][i] == board[1][i] && board[0][i] == board[2][i] {
			/*
			   | | | |      | | | |       |X|X|X|
			   | | | |  OR  |X|X|X|  OR   | | | |
			   |X|X|X|      | | | |       | | | |
			*/
			return board[0][i]
		}
	}

	// Check diagonals
	if board[0][0] != EMPTY && board[0][0] == board[1][1] && board[0][0] == board[2][2] {
		/*
		   | | |X|
		   | |X| |
		   |X| | |
		*/
		return board[0][0]
	} else if board[2][0] != EMPTY && board[2][0] == board[1][1] && board[2][0] == board[0][2] {
		/*
		   |X| | |
		   | |X| |
		   | | |X|
		*/
		return board[2][0]
	}

	return EMPTY // Winning conditions not satisfied for either player.
}

func (board *TictactoeBoard) Clear() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = EMPTY
		}
	}
}

// ValidMoves returns a slice of *Move, containing all the "legal"
// moves that can still be made on the Tic-tac-toe board in question.
func (board *TictactoeBoard) ValidMoves(boardX, boardY int) []*Move {
	validMoves := make([]*Move, 0, 9) // Pre-allocate capacity for up to 9 moves (the max)
	if board.HasWinner() != EMPTY {
		// If the board has already been won no moves can be made on it.
		return nil
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == EMPTY {
				validMoves = append(validMoves, &Move{boardX, boardY, i, j})
			}
		}
	}

	return validMoves
}

// HasWinner uses similar logic as in TictactoeBoard:HasWinner to
// check whether either player has already won the ultimate board.
func (board *UltimateBoard) HasWinner() int {
	for i := 0; i < 3; i++ {
		if board[i][0].HasWinner() != EMPTY && board[i][0].HasWinner() == board[i][1].HasWinner() && board[i][0].HasWinner() == board[i][2].HasWinner() {
			/*
			   |X| | |      | |X| |       | | |X|
			   |X| | |  OR  | |X| |  OR   | | |X|
			   |X| | |      | |X| |       | | |X|
			*/
			return board[i][0].HasWinner()
		}

		if board[0][i].HasWinner() != EMPTY && board[0][i].HasWinner() == board[1][i].HasWinner() && board[0][i].HasWinner() == board[2][i].HasWinner() {
			/*
			   | | | |      | | | |       |X|X|X|
			   | | | |  OR  |X|X|X|  OR   | | | |
			   |X|X|X|      | | | |       | | | |
			*/
			return board[0][i].HasWinner()
		}
	}

	// Check diagonals
	if board[0][0].HasWinner() != EMPTY && board[0][0].HasWinner() == board[1][1].HasWinner() && board[0][0].HasWinner() == board[2][2].HasWinner() {
		/*
		   | | |X|
		   | |X| |
		   |X| | |
		*/
		return board[0][0].HasWinner()
	} else if board[2][0].HasWinner() != EMPTY && board[2][0].HasWinner() == board[1][1].HasWinner() && board[2][0].HasWinner() == board[0][2].HasWinner() {
		/*
		   |X| | |
		   | |X| |
		   | | |X|
		*/
		return board[2][0].HasWinner()
	}

	return EMPTY // Winning conditions not satisfied for either player.
}

// Clear clears the UltimateBoard, setting every
// square to Empty.
func (board *UltimateBoard) Clear() {
	var emptyBoard TictactoeBoard
	emptyBoard.Clear()

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = emptyBoard
		}
	}
}

// ValidMoves returns a slice of *Move, containing all the "legal"
// moves that can still be made on the Tic-tac-toe board in question.
func (board *UltimateBoard) ValidMoves() []*Move {
	validMoves := make([]*Move, 81) // Pre-allocate capacity for up to 81 moves (the max)
	movesNum := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j].HasWinner() != EMPTY {
				// If the board is already won there are no valid moves left.
				continue
			}

			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					if board[i][j][k][l] == EMPTY {

						validMoves[movesNum] = &Move{i, j, k, l}
						movesNum += 1
					}
				}
			}
		}
	}

	return validMoves[0:movesNum]
}

// AllPossibleMoves returns a slice of *Move, containing all the
// moves that can still be made on the Tic-tac-toe board in
// question, including moves on already won / lost boards.
func (board *UltimateBoard) AllPossibleMoves() []*Move {
	validMoves := make([]*Move, 81) // Pre-allocate capacity for up to 81 moves (the max)
	movesNum := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					if board[i][j][k][l] == EMPTY {

						validMoves[movesNum] = &Move{i, j, k, l}
						movesNum += 1
					}
				}
			}
		}
	}

	return validMoves[0:movesNum]
}

// Copy returns a pointer to a copy of the UltimateBoard
func (board *UltimateBoard) Copy() *UltimateBoard {
	originalBoard := *board
	boardCopy := originalBoard
	return &boardCopy
}
