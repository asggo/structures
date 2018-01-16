// Package merkle generates a merkle tree from a list of byte slices. Each
// byte slice (block) is hashed using SHA256 and the resulting hashes are
// used to build the Merkle tree.
package merkle

import (
	"bytes"
	"crypto/sha256"
)

// The Merkle type represents a binary Merkle tree.
type Merkle struct {
	digest   [32]byte
	children [2]*Merkle
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
		m.children[1].Diff(m2.children[1], diffs)
	}
}


// Recursively build a Merkle tree using the slice of byte slices.
func NewMerkle(blocks [][]byte) *Merkle {
	var data []byte
	m := new(Merkle)

	switch len(blocks) {
	case 1:
		m.children[0] = &Merkle{digest: sha256.Sum256(blocks[0])}
		data = m.children[0].digest[:]
	case 2:
		m.children[0] = &Merkle{digest: sha256.Sum256(blocks[0])}
		m.children[1] = &Merkle{digest: sha256.Sum256(blocks[1])}
		data = append(m.children[0].digest[:], m.children[1].digest[:]...)

	default:
		half := len(blocks) / 2

		m.children[0] = NewMerkle(blocks[:half])
		m.children[1] = NewMerkle(blocks[half:])
		data = append(m.children[0].digest[:], m.children[1].digest[:]...)
	}

	m.digest = sha256.Sum256(data)

	return m
}
