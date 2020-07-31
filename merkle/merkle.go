// Package merkle generates a merkle tree from a list of byte slices. Each
// byte slice (block) is hashed using SHA256 and the resulting hashes are
// used to build the Merkle tree.
package merkle

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// The Merkle type represents a binary Merkle tree.
type Merkle struct {
	digest  [32]byte
	encoded string
	left    *Merkle
	right   *Merkle
	level   int
	nodes   int
}

func (m *Merkle) String() string {
	return fmt.Sprintf("Depth: %d Nodes: %d Digest: %s", m.level, m.nodes, m.encoded)
}

// Equal returns true if two Merkle trees are equivalent
func (m *Merkle) Equal(m2 *Merkle) bool {
	return m.encoded == m2.encoded
}

// Digest returns the hex encoded digest of the Merkle tree
func (m *Merkle) Digest() string {
	return m.encoded
}

// Diff returns a slice of encoded digests from m2 which are different from
// those in m1.
func (m *Merkle) Diff(m2 *Merkle, diffs *[]string) {

	if m.left != nil {
		m.left.Diff(m2.left, diffs)
	}

	if m.right != nil {
		m.right.Diff(m2.right, diffs)
	}

	// We are in a leaf node. Do our comparison.
	if m.left == nil && m.encoded != m2.encoded {
		*diffs = append(*diffs, m2.encoded)
	}
}

// newLeafNode returns a new leaf node in the Merkle tree created from the
// given block.
func newLeafNode(block []byte) *Merkle {
	m := new(Merkle)

	m.digest = sha256.Sum256(block)
	m.encoded = hex.EncodeToString(m.digest[:])
	m.level = 0
	m.nodes = 1

	return m
}

// newMerkleNode returns a new merkle node created from the give leaf nodes.
func newMerkleNode(leaf1, leaf2 *Merkle) *Merkle {
	m := new(Merkle)

	if leaf2 == nil {
		m.digest = leaf1.digest
		m.nodes = leaf1.nodes
	} else {
		m.digest = sha256.Sum256(append(leaf1.digest[:], leaf2.digest[:]...))
		m.nodes = leaf1.nodes + leaf2.nodes
	}

	m.encoded = hex.EncodeToString(m.digest[:])
	m.left = leaf1
	m.right = leaf2
	m.level = leaf1.level + 1

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
