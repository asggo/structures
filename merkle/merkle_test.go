package merkle

import (
	"fmt"
	"testing"
)

func TestMerkle(t *testing.T) {
	b1 := []byte("aaaaa")
	b2 := []byte("bbbbb")
	b3 := []byte("ccccc")
	b4 := []byte("ddddd")
	b5 := []byte("eeeee")

	blocks0 := [][]byte{b1, b2, b3}
	blocks1 := [][]byte{b1, b2, b3, b4, b5}
	blocks2 := [][]byte{b5, b4, b3, b2, b1}

	m0 := NewMerkle(blocks0)
	fmt.Println(m0.String())
	m1 := NewMerkle(blocks1)
	m2 := NewMerkle(blocks1)
	m3 := NewMerkle(blocks2)

	if !m1.Equal(m2) {
		t.Error("Merkle trees should be equal.")
	}

	if m1.Equal(m3) {
		t.Error("Merkle trees should not be equal.")
	}

	var diffs [][32]byte

	m1.Diff(m2, &diffs)

	if len(diffs) != 0 {
		t.Error("Differences found in identical trees.")
	}

	diffs = nil
	m1.Diff(m3, &diffs)

	if len(diffs) != 4 {
		t.Error("Expected ", 4, "got", len(diffs))
	}
}
