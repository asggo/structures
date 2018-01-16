// The quadtree package implements a QuadTree structure that can be used for
// any application requiring the identification of objects that are physically
// close to one another on a Cartesian plane. The QuadTree can store any
// object that satisfies the Boxer interface.
package quadtree

import (
	"fmt"
)

// The maximum number of values in any one node.
const maxNodeSize = 16

// Boxer is the interface for the Box method.
type Boxer interface {
	Box() *Box
}

// Node defines a single node in the Quadtree.
type Node struct {
	level       int
	children    [4]*Node
	values      []Boxer
	boundingBox *Box
}

// String returns a string representing the Node structure.
func (n *Node) String() string {
	s := fmt.Sprintf("Level: %d\nValues: %d\nChildren: %t\nBox: %s\n",
		n.level, len(n.values), n.children[0] != nil, n.boundingBox.String())

	if n.children[0] != nil {
		for i, _ := range n.children {
			s = s + n.children[i].String()
		}
	}

	return s
}

// Count recursively counts all of the values stored in the Quadtree. The
// value is stored in the given int pointer.
func (n *Node) Count(count *int) {
	*count = *count + len(n.values)

	if n.children[0] != nil {
		for i, _ := range n.children {
			n.children[i].Count(count)
		}
	}
}

// NodeCount recursively counts all of the nodes in the Quadtree. The value
// is stored in the given int pointer.
func (n *Node) NodeCount(count *int) {
	*count = *count + 1

	if n.children[0] != nil {
		for i, _ := range n.children {
			n.children[i].NodeCount(count)
		}
	}
}

// Clear recursively clears each node from the Quadtree.
func (n *Node) Clear() {
	n.values = nil

	if n.children[0] != nil {
		for i, _ := range n.children {
			n.children[i].Clear()
		}
	}
}

// Split creates four new child nodes and reinserts each value so that it ends
// up in the appropriate child node or back in the parent node.
func (n *Node) split() {
	quads := n.boundingBox.Quarter()

	n.children[0] = NewNode(n.level+1, quads[0])
	n.children[1] = NewNode(n.level+1, quads[1])
	n.children[2] = NewNode(n.level+1, quads[2])
	n.children[3] = NewNode(n.level+1, quads[3])

	// Make a copy of our values
	var values []Boxer
	values = append(values, n.values...)

	// Clear out the current values
	n.values = nil

	// Reinsert our values
	for i, _ := range values {
		n.Insert(values[i])
	}
}

// Index returns the index value of the child, if any, that contains the given
// value.
func (n *Node) index(v Boxer) int {
	if n.children[0] == nil {
		return -1
	}

	for i, _ := range n.children {
		if n.children[i].boundingBox.Contains(v.Box()) {
			return i
		}
	}

	return -1
}

// Insert adds a new value to the appropriate child node. If there are no
// children or if the value does not fit into one of the children, the value
// is added to this node. Split the node if it is full.
func (n *Node) Insert(v Boxer) {
	// If this node does not contain the given box return.
	if !n.boundingBox.Contains(v.Box()) {
		return
	}

	i := n.index(v)

	if i == -1 {
		n.values = append(n.values, v)
		if len(n.values) > maxNodeSize {
			n.split()
		}
	} else {
		n.children[i].Insert(v)
	}
}

// Retrieve returns all values that intersect with the given box. The values
// are appended to the given Boxer slice pointer.
func (n *Node) Retrieve(b *Box, values *[]Boxer) {
	// If this node does not intersect with the given box return.
	if !n.boundingBox.Intersects(b) {
		return
	}

	// Find all values in this node that intersect with the given box.
	for i, _ := range n.values {
		if b.Intersects(n.values[i].Box()) {
			*values = append(*values, n.values[i])
		}
	}

	// Recurse into each child node.
	if n.children[0] != nil {
		for i, _ := range n.children {
			n.children[i].Retrieve(b, values)
		}
	}
}

// NewNode creates a new Quadtree node.
func NewNode(level int, box *Box) *Node {
	n := new(Node)

	n.level = level
	n.boundingBox = box

	return n
}
