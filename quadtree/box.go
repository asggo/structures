package quadtree

import (
	"fmt"
)

// Box defines an axis-aligned bounding box using the (x,y) value of the
// bottom, left corner along with the width and height of the box.
type Box struct {
	x      int
	y      int
	width  int
	height int
}

func (b *Box) String() string {
	return fmt.Sprintf("(%d, %d), (%d, %d)", b.Left(), b.Bottom(), b.Right(), b.Top())
}

// Left returns the x-value of the left side of the box.
func (b *Box) Left() int { return b.x }

// Right returns the x-value of the right side of the box.
func (b *Box) Right() int { return b.x + b.width }

// Top returns the y-value of the top side of the box.
func (b *Box) Top() int { return b.y + b.height }

// Bottom returns the y-value of the bottom side of the box.
func (b *Box) Bottom() int { return b.y }

// CenterX returns the x value of the center of the box
func (b *Box) CenterX() int {
	return b.x + int(b.width/2)
}

// CenterY returns the y value of the center of the box
func (b *Box) CenterY() int {
	return b.y + int(b.height/2)
}

// Quarter splits a box into its four quadrants starting at the top right
// quadrant and going counter-clockwise.
func (b *Box) Quarter() [4]*Box {
	var quarters [4]*Box

	w := int(b.width / 2)
	h := int(b.height / 2)

	quarters[0] = NewBox(b.Left()+w, b.Bottom()+h, w, h) // Top Right
	quarters[1] = NewBox(b.Left(), b.Bottom()+h, w, h)   // Top Left
	quarters[2] = NewBox(b.Left(), b.Bottom(), w, h)     // Bottom Left
	quarters[3] = NewBox(b.Left()+w, b.Bottom(), w, h)   // Bottom Right

	return quarters
}

// Contains returns true if the given box is fully contained by this box.
func (b *Box) Contains(c *Box) bool {
	x := (b.Left() <= c.Left()) && (b.Right() >= c.Right())
	y := (b.Top() >= c.Top()) && (b.Bottom() <= c.Bottom())

	return x && y
}

// ContainsCenter returns true if the center of the given box is contained by
// this box.
func (b *Box) ContainsCenter(c *Box) bool {
	x := (b.Left() <= c.CenterX()) && (b.Right() >= c.CenterX())
	y := (b.Top() >= c.CenterY()) && (b.Bottom() <= c.CenterY())

	return x && y
}

// Intersects returns true if the give box overlaps this box.
func (b *Box) Intersects(c *Box) bool {
	xIntersect := (b.Left() < c.Right()) && (c.Left() < b.Right())
	yIntersect := (b.Top() > c.Bottom()) && (c.Top() > b.Bottom())

	return xIntersect && yIntersect
}

// NewBox creates a new Box structure.
func NewBox(x, y, width, height int) *Box {
	b := new(Box)

	b.x = x
	b.y = y
	b.width = width
	b.height = height

	return b
}
