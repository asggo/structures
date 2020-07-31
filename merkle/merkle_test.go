package merkle

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"
)

func sum(block []byte) []byte {
	digest := sha256.Sum256(block)

	return digest[:]
}

func encode(data []byte) string {
	return hex.EncodeToString(data)
}

func TestMerkle(t *testing.T) {
	b1 := []byte("aaaaa")
	b2 := []byte("bbbbb")
	b3 := []byte("ccccc")
	b4 := []byte("ddddd")
	b5 := []byte("eeeee")

	// Verify that hashing works as expected.

	// The Merkle tree has one first level node consisting of (b1, nil). The
	// digest will be: sum(b1, nil)
	//  node1
	//    |
	//   b1
	//
	fmt.Println("Testing 1 block Merkle Tree")
	m1 := NewMerkle([][]byte{b1})
	node1 := sum(b1)
	if m1.Digest() != encode(node1) {
		t.Fatalf("Expected digest %s, received %s.", encode(node1), m1.Digest())
	}

	// The Merkle tree has one first level node consisting of (b1, b2). The
	// digest will be: sum(sum(b1), sum(b2)).
	//
	//  node1
	//   / \
	// b1   b2
	//
	fmt.Println("Testing 2 block Merkle Tree")
	m2 := NewMerkle([][]byte{b1, b2})
	node1 = sum(append(sum(b1), sum(b2)...))
	if m2.Digest() != encode(node1) {
		t.Fatalf("Expected digest %s, received %s.", encode(node1), m2.Digest())
	}

	// The Merkle tree has two first level nodes. The first node is (b1, b2)
	// and the second node is (b3, nil). The digest will be
	// sum(sum(b1, b2), sum(b3)).
	//
	//      node3
	//       /  \
	//  node1    node2
	//   / \      / \
	// b1   b2  b3   nil
	//
	fmt.Println("Testing 3 block Merkle Tree")
	m3 := NewMerkle([][]byte{b1, b2, b3})
	node1 = sum(append(sum(b1), sum(b2)...))
	node2 := sum(b3)
	node3 := sum(append(node1, node2...))
	if m3.Digest() != encode(node3) {
		t.Fatalf("Expected digest %s, received %s.", encode(node3), m3.Digest())
	}

	// The Merkle tree has two first level nodes. The first node is (b1, b2)
	// and the second node is (b3, b4). The digest will be
	// sum(sum(b1, b2), sum(b3, b4)).
	//
	//      node3
	//       /  \
	//  node1    node2
	//   / \      / \
	// b1   b2  b3   b4
	//
	fmt.Println("Testing 4 block Merkle Tree")
	m4 := NewMerkle([][]byte{b1, b2, b3, b4})
	node1 = sum(append(sum(b1), sum(b2)...))
	node2 = sum(append(sum(b3), sum(b4)...))
	node3 = sum(append(node1, node2...))
	if m4.Digest() != encode(node3) {
		t.Fatalf("Expected digest %s, received %s.", encode(node3), m4.Digest())
	}

	// The Merkle tree has two first level nodes, the first is (b1, b2) and
	// the second is (b3, b4). It has two second level nodes, the first is
	// node3 and the second is (b5, nil). The digest will be
	// sum(sum(sum(b1, b2), sum(b3, b4)), sum(b5, nil))
	//
	//            node5
	//           /     \
	//      node3       node4
	//       /  \        |\
	//  node1    node2   | \
	//   / \      / \    |  \
	// b1   b2  b3   b4  b5  nil
	//
	fmt.Println("Testing 5 block Merkle Tree")
	m5 := NewMerkle([][]byte{b1, b2, b3, b4, b5})
	node1 = sum(append(sum(b1), sum(b2)...))
	node2 = sum(append(sum(b3), sum(b4)...))
	node3 = sum(append(node1, node2...))
	node4 := sum(b5)
	node5 := sum(append(node3, node4...))
	if m5.Digest() != encode(node5) {
		t.Fatalf("Expected digest %s, received %s.", encode(node5), m5.Digest())
	}

	// Test Equality and Diff checks
	blocks1 := [][]byte{b1, b2, b3, b4, b5}
	blocks2 := [][]byte{b5, b4, b3, b2, b1}

	m1 = NewMerkle(blocks1)
	m2 = NewMerkle(blocks1)
	m3 = NewMerkle(blocks2)

	fmt.Println("Testing equality check")
	if !m1.Equal(m2) {
		t.Error("Merkle trees should be equal.")
	}

	if m1.Equal(m3) {
		t.Error("Merkle trees should not be equal.")
	}

	fmt.Println("Testing diff check")
	var diffs []string

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

func benchmarkMerkle(size int, b *testing.B) {
	blocks := make([][]byte, size)
	for i := range blocks {
		blocks[i] = []byte("aaaaa")
	}

	//b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewMerkle(blocks)
	}
}

func BenchmarkMerkle10(b *testing.B) {
	benchmarkMerkle(10, b)
}

func BenchmarkMerkle100(b *testing.B) {
	benchmarkMerkle(100, b)
}

func BenchmarkMerkle1000(b *testing.B) {
	benchmarkMerkle(1000, b)
}

func BenchmarkMerkle10000(b *testing.B) {
	benchmarkMerkle(10000, b)
}

func BenchmarkMerkle100000(b *testing.B) {
	benchmarkMerkle(100000, b)
}

func BenchmarkMerkle1000000(b *testing.B) {
	benchmarkMerkle(1000000, b)
}
