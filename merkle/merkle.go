// Package merkle generates a merkle tree from a list of byte slices. Each
// byte slice (block) is hashed using SHA256 and the resulting hashes are
// used to build the Merkle tree.
package merkle

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// The Merkle type represents a binary Merkle tree.
type Merkle struct {
	digest   [32]byte
	encoded  string
	children [2]*Merkle
}

func (m *Merkle) String() string {
	var s string

	s = fmt.Sprintf("%s\n", m.encoded)

	if m.children[0] != nil {
		s = s + fmt.Sprintf("  %s", m.children[0].String())
	}

	if m.children[1] != nil {
		s = s + fmt.Sprintf("  %s", m.children[1].String())
	}

	return s
}

// Equal returns true if two Merkle trees are equivalent
func (m *Merkle) Equal(m2 *Merkle) bool {
	return bytes.Equal(m.digest[:], m2.digest[:])
}

// Diff returns a slice of digests from m2 which are different from those in m1.
func (m *Merkle) Diff(m2 *Merkle, diffs *[][32]byte) {
	if m.Equal(m2) {
		return
	}

	if m.children[0] == nil {
		*diffs = append(*diffs, m2.digest)
	} else {
		m.children[0].Diff(m2.children[0], diffs)
		if m.children[1] != nil {
			m.children[1].Diff(m2.children[1], diffs)
		}
	}
}

// newLeafNode returns a new leaf node in the Merkle tree created from the
// given block.
func newLeafNode(block []byte) *Merkle {
	m := new(Merkle)

	m.digest = sha256.Sum256(block)
	m.encoded = hex.EncodeToString(m.digest[:])

	return m
}

// newMerkleNode returns a new merkle node created from the give leaf nodes.
func newMerkleNode(leaf1, leaf2 *Merkle) *Merkle {
	m := new(Merkle)

	if leaf2 == nil {
		m.digest = leaf1.digest
	} else {
		m.digest = sha256.Sum256(append(leaf1.digest[:], leaf2.digest[:]...))
	}

	m.encoded = hex.EncodeToString(m.digest[:])
	m.children[0] = leaf1
	m.children[1] = leaf2

	return m
}

// Build a Merkle tree using the slice of byte slices.
func NewMerkle(blocks [][]byte) *Merkle {
	var leaves []*Merkle

	// Build our leaf nodes
	for i, _ := range blocks {
		leaves = append(leaves, newLeafNode(blocks[i]))
	}

	// Build parent nodes until there is only one parent.
	for {
		if len(leaves) == 1 {
			break
		}

		var newLeaves []*Merkle

		// Create new nodes from pairs of nodes
		for i := 0; i < len(leaves)-1; i += 2 {
			newLeaves = append(newLeaves, newMerkleNode(leaves[i], leaves[i+1]))
		}

		// Append the remaining node when there are an uneven number.
		if len(leaves)%2 != 0 {
			newLeaves = append(newLeaves, newMerkleNode(leaves[len(leaves)-1], nil))
		}

		leaves = nil
		leaves = append(leaves, newLeaves...)
	}

	return leaves[0]
}
