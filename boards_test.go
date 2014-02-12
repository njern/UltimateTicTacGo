package main

import "testing"

func TestTicTacToeBordHasWinner(t *testing.T) {
	var board TictactoeBoard
	board.Clear()

	if board.HasWinner() != EMPTY {
		t.Error("An empty board should not have a winner!")
	}

	// Check vertical "wins" for player 1
	for i := 0; i < 3; i++ {
		board.Clear()
		board[i][0] = PLAYER_1_CONTROLLED
		board[i][1] = PLAYER_1_CONTROLLED
		board[i][2] = PLAYER_1_CONTROLLED
		if board.HasWinner() != PLAYER_1_CONTROLLED {
			t.Error("With three X's in a (vertical) row, player 1 should be the winner!")
		}
	}

	// Check horizontal "wins" for player 2
	for i := 0; i < 3; i++ {
		board.Clear()
		board[0][i] = PLAYER_2_CONTROLLED
		board[1][i] = PLAYER_2_CONTROLLED
		board[2][i] = PLAYER_2_CONTROLLED
		if board.HasWinner() != PLAYER_2_CONTROLLED {
			t.Error("With three O's in a (horizontal) row, player 2 should be the winner!")
		}
	}

	// Check diagonal 1 (bottom left to top right) for player 1
	board.Clear()
	board[0][0] = PLAYER_1_CONTROLLED
	board[1][1] = PLAYER_1_CONTROLLED
	board[2][2] = PLAYER_1_CONTROLLED

	if board.HasWinner() != PLAYER_1_CONTROLLED {
		t.Error("With three X's in a (diagonal) row, player 1 should be the winner!")
	}

	// Check diagonal 2 (top left to bottom right) for player 2
	board.Clear()
	board[2][0] = PLAYER_2_CONTROLLED
	board[1][1] = PLAYER_2_CONTROLLED
	board[0][2] = PLAYER_2_CONTROLLED

	if board.HasWinner() != PLAYER_2_CONTROLLED {
		t.Error("With three X's in a (diagonal) row, player 2 should be the winner!")
	}

	// Try a few boards which should be undecided.
	for i := 0; i < 3; i++ {
		board.Clear()
		board[i][0] = PLAYER_1_CONTROLLED
		board[i][1] = PLAYER_2_CONTROLLED
		board[i][2] = PLAYER_1_CONTROLLED
		if board.HasWinner() != EMPTY {
			t.Error("Player 2 blocked the middle tile - there should be no winner here!")
		}

		board.Clear()
		board[0][i] = PLAYER_2_CONTROLLED
		board[1][i] = PLAYER_1_CONTROLLED
		board[2][i] = PLAYER_2_CONTROLLED
		if board.HasWinner() != EMPTY {
			t.Error("Player 1 blocked the middle tile - there should be no winner here!")
		}
	}
}

func TestUltimateBoardHasWinner(t *testing.T) {
	var board UltimateBoard
	board.Clear()

	var player1WonBoard TictactoeBoard
	player1WonBoard.Clear()
	player1WonBoard[0][0] = PLAYER_1_CONTROLLED
	player1WonBoard[1][1] = PLAYER_1_CONTROLLED
	player1WonBoard[2][2] = PLAYER_1_CONTROLLED

	var player2WonBoard TictactoeBoard
	player2WonBoard.Clear()
	player2WonBoard[2][0] = PLAYER_2_CONTROLLED
	player2WonBoard[1][1] = PLAYER_2_CONTROLLED
	player2WonBoard[0][2] = PLAYER_2_CONTROLLED

	if board.HasWinner() != EMPTY {
		t.Error("An empty board should not have a winner!")
	}

	// Check vertical "wins" for player 1
	for i := 0; i < 3; i++ {
		board.Clear()
		board[i][0] = player1WonBoard
		board[i][1] = player1WonBoard
		board[i][2] = player1WonBoard
		if board.HasWinner() != PLAYER_1_CONTROLLED {
			t.Error("With three X's in a (vertical) row, player 1 should be the winner!")
		}
	}

	// Check horizontal "wins" for player 2
	for i := 0; i < 3; i++ {
		board.Clear()
		board[0][i] = player2WonBoard
		board[1][i] = player2WonBoard
		board[2][i] = player2WonBoard
		if board.HasWinner() != PLAYER_2_CONTROLLED {
			t.Error("With three O's in a (horizontal) row, player 2 should be the winner!")
		}
	}

	// Check diagonal 1 (bottom left to top right) for player 1
	board.Clear()
	board[0][0] = player1WonBoard
	board[1][1] = player1WonBoard
	board[2][2] = player1WonBoard

	if board.HasWinner() != PLAYER_1_CONTROLLED {
		t.Error("With three X's in a (diagonal) row, player 1 should be the winner!")
	}

	// Check diagonal 2 (top left to bottom right) for player 2
	board.Clear()
	board[2][0] = player2WonBoard
	board[1][1] = player2WonBoard
	board[0][2] = player2WonBoard

	if board.HasWinner() != PLAYER_2_CONTROLLED {
		t.Error("With three O's in a (diagonal) row, player 2 should be the winner!")
	}

	// Try a few boards which should be undecided.
	for i := 0; i < 3; i++ {
		board.Clear()
		board[i][0] = player1WonBoard
		board[i][1] = player2WonBoard
		board[i][2] = player1WonBoard
		if board.HasWinner() != EMPTY {
			t.Error("Player 2 blocked the middle tile - there should be no winner here!")
		}

		board.Clear()
		board[0][i] = player2WonBoard
		board[1][i] = player1WonBoard
		board[2][i] = player2WonBoard
		if board.HasWinner() != EMPTY {
			t.Error("Player 1 blocked the middle tile - there should be no winner here!")
		}
	}
}

func TestUltimateBoardCopy(t *testing.T) {
	var board UltimateBoard
	board.Clear()

	anotherBoard := board.Copy()

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] != anotherBoard[i][j] {
				t.Error("Failed to Copy() the UltimateBoard!")
			}
		}
	}
}

func TestUltimateBoardAllPossibleMoves(t *testing.T) {
	var board UltimateBoard
	board.Clear()

	possibleMoves := board.AllPossibleMoves()
	if len(possibleMoves) != 81 {
		t.Error("An empty board should have 81 possible moves!")
	}
}

func BenchmarkHasWinnerEmptyBoard(b *testing.B) {
	var board TictactoeBoard

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = board.HasWinner()
	}
}

func BenchmarkHasWinnerEmptyUltimateBoard(b *testing.B) {
	var board UltimateBoard

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = board.HasWinner()
	}
}

func BenchmarkValidMovesEmptyBoard(b *testing.B) {
	var board TictactoeBoard

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = board.ValidMoves(1, 1)
	}
}

func BenchmarkValidMovesEmptyUltimateBoard(b *testing.B) {
	var board UltimateBoard

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = board.ValidMoves()
	}
}

func BenchmarkMoveCreate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Move{1, 1, 1, 1}
	}
}
