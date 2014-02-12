package main

// Move describes a move made by a player on the UltimateBoard.
type Move struct {
	BoardX int
	BoardY int
	TileX  int
	TileY  int
}

// Copy returns a pointer to a copy of Move.
func (m *Move) Copy() *Move {
	moveCopy := Move{m.BoardX, m.BoardY, m.TileX, m.TileY}
	return &moveCopy
}
