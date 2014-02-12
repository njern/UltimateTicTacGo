package main

import "testing"

func TestRandomBot(t *testing.T) {
	var board UltimateBoard
	board.Clear()
	firstMove := Move{-1, -1, -1, -1}

	// Test playing the first move
	randomMove := RandomBot(&firstMove, &board)
	if randomMove == nil {
		t.Error("RandomBot failed to produce a random move when making the first move!")
	}

	// Test playing another move, see that RandomBot follows the rules.
	anotherMove := Move{1, 1, 1, 1}
	randomMove = RandomBot(&anotherMove, &board)
	if randomMove.BoardX != 1 || randomMove.BoardY != 1 {
		t.Error("RandomBot did not stick to the board it was forced to!")
		t.Error("Instead of (1,1), it played on board (", randomMove.BoardX, ",", randomMove.BoardY, ")")
	}
}

func TestMonteCarloBot(t *testing.T) {
	var board UltimateBoard
	board.Clear()
	firstMove := Move{-1, -1, -1, -1}

	// Test playing the first move
	smartMove := MonteCarloBot(1, &firstMove, &board)
	if smartMove == nil {
		t.Error("MonteCarloBot failed to produce a random move when making the first move!")
	}

	// Test playing another move, see that MonteCarloBot follows the rules.
	anotherMove := Move{1, 1, 1, 1}
	smartMove = MonteCarloBot(2, &anotherMove, &board)
	if smartMove.BoardX != 1 || smartMove.BoardY != 1 {
		t.Error("MonteCarloBot did not stick to the board it was forced to!")
		t.Error("Instead of (1,1), it played on board (", smartMove.BoardX, ",", smartMove.BoardY, ")")
	}
}
