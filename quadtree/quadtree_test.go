package quadtree

import (
	"fmt"
	"math/rand"
	"testing"
)

type Object struct {
	x int
	y int
	r int
}

func (o *Object) Box() *Box {
	return NewBox(o.x-o.r, o.y-o.r, o.r*2, o.r*2)
}

func (o *Object) String() string {
	return fmt.Sprintf("[(%d, %d), %d]", o.x, o.y, o.r)
}

func randomObject(box *Box) *Object {
	o := new(Object)
	o.r = 4

	xRange := (box.Right() - box.Left()) - (4 * o.r)
	yRange := (box.Top() - box.Bottom()) - (4 * o.r)

	o.x = int(rand.Intn(xRange)) + o.r + box.Left()
	o.y = int(rand.Intn(yRange)) + o.r + box.Bottom()

	return o
}

func TestQuadTree(t *testing.T) {
	box := NewBox(-32, -32, 64, 64)
	qt := NewNode(0, box)

	// Generate an object in the box and one not in the box. Insert them both
	// and retrieve all objects in the quadtree. Verify only one object is
	// retrieved.
	o1 := &Object{x: -3, y: 3, r: 3}
	o2 := &Object{x: 62, y: 62, r: 4}

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

	// Generate 100 random objects and insert them into the quadtree. Verify
	// all 100 objects are inserted into the quadtree by counting the number of
	// values in the tree.
	fmt.Println("")
	for i := 0; i < 100; i++ {
		qt.Insert(randomObject(box))
	}

	var valCount int

	qt.Count(&valCount)

	if valCount != 100 {
		t.Error("Expected ", 100, "got", valCount)
	}
}
