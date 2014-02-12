package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	EMPTY               = 45  // '-' - represents an empty square
	PLAYER_1_CONTROLLED = 88  // 'X' - a square controlled by player 1
	PLAYER_2_CONTROLLED = 79  // 'O' - a square controlled by player 2
	TIME_TO_THINK       = 5.9 // How long the Monte Carlo bot can think before making it's move (seconds)
)

// main, in this case, reads in the board state from HackerRank
// and emits the next Move as a space separated string.
func main() {
	reader := bufio.NewReader(os.Stdin)
	var playerNumber int
	var lastMove *Move
	var board UltimateBoard
	var rowIndex int

	for {
		line, err := reader.ReadString('\n')
		lineItems := strings.Fields(line)

		if playerNumber == 0 {
			// First line, we're reading in which player we're playing as.
			if lineItems[0] == "X" {
				playerNumber = 1
			} else {
				playerNumber = 2
			}
		} else if lastMove == nil {
			// Second line, reading in which board we get to play on next
			// -1, -1 means we can play on any board.
			lastMoveTileX, _ := strconv.Atoi(lineItems[0])
			lastMoveTileY, _ := strconv.Atoi(lineItems[1])
			lastMove = &Move{0, 0, lastMoveTileX, lastMoveTileY}
		} else {
			// Now reading in the lines representing the
			// current board state.
			row := lineItems[0]
			for index, cell := range row {
				board[rowIndex/3][index/3][rowIndex%3][index%3] = int(cell)
			}
			rowIndex += 1
		}

		if err != nil {
			// This is the last line -> print the bot's next move in HackerRank's preferred format.
			move := MonteCarloBot(playerNumber, lastMove, &board)
			fmt.Printf("%d %d %d %d\n", move.BoardX, move.BoardY, move.TileX, move.TileY)
			break
		}
	}
}
