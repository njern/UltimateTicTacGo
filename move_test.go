package main

import "testing"

func TestCopyMove(t *testing.T) {
	m := Move{1, 2, 3, 4}
	m2 := m.Copy()

	if m.BoardX != m2.BoardX || m.BoardY != m2.BoardY || m.TileX != m2.TileX || m.TileY != m2.TileY {
		t.Error("Move::Copy() failed to produce a valid copy.")
	}
}
