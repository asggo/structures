package quadtree

import (
    "testing"
    "fmt"
)

type Object struct {
    x   int
    y   int
    r   int
}

func (o *Object) Box() *Box {
    return NewBox(o.x-o.r, o.y+o.r, o.r*2, o.r*2)
}

func (o *Object) String() string {
    return fmt.Sprintf("[(%d, %d), %d]", o.x, o.y, o.r)
}

func NewObject(x, y, r int) *Object {
    o := new(Object)

    o.x = x
    o.y = y
    o.r = r

    return o
}

func TestQuadTree(t *testing.T) {
    box := NewBox(-32, 32, 32, 32)
    qt := NewNode(0, box)

    // Generate an object in the box and one not in the box. Insert them both
    // and retrieve all objects in the quadtree. Verify only one object is
    // retrieved.
    o1 := NewObject(-3, 3, 3)
    o2 := NewObject(0, 0, 4)

    qt.Insert(o1)
    qt.Insert(o2)

    var objects []Boxer
    qt.Retrieve(box, &objects)

    if len(objects) != 1 {
        t.Error("Expected ", 1, "got", len(objects))
    }

    // Clear the tree and verify there are no values stored in the tree.
    var count int

    qt.Clear()
    qt.Count(&count)

    if count != 0 {
        t.Error("Expected ", 0, "got", count)
    }

    // Generate maxNodeSize + 1 objects and insert them into the quad tree.
    // Verify the Quadtree successfully split by verifying the number of
    // values stored in the tree and the number of nodes in the tree.
    for i:=0; i<= maxNodeSize; i++ {
        o := NewObject(-(i+5), i+5, 2)
        qt.Insert(o)
    }

    var valCount int
    var nodeCount int

    qt.Count(&valCount)
    qt.NodeCount(&nodeCount)

    if valCount != maxNodeSize + 1 {
        t.Error("Expected ", maxNodeSize + 1, "got", valCount)
    }

    if nodeCount != 5 {
        t.Error("Expected ", 5, "nodes got", nodeCount)
    }
}
