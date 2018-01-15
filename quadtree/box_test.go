package quadtree

import (
    "testing"
)


func TestBox(t *testing.T) {
    b1 := NewBox(0, 10, 100, 100)  // (0, 10), (100, -90)
    b2 := NewBox(10, 5, 40, 35)    // (10, 5), (50, -30)
    b3 := NewBox(-10, -10, 5, 10)  // (-10, -10), (-5, 0)
    b4 := NewBox(-10, 0, 10, 10)   // (-10, 0), (0, -10)

    if !b4.Contains(b4) {
        t.Error("Box should contain itself.")
    }

    if !b1.Contains(b2) {
        t.Errorf("%s does not contain %s\n", b1.String(), b2.String())
    }

    if b1.Contains(b3) {
        t.Errorf("%s contains %s\n", b1.String(), b3.String())
    }

    if !b4.Intersects(b4) {
        t.Error("Box should intersect itself.")
    }

    if !b1.Intersects(b2) {
        t.Errorf("%s does not intersect %s\n", b1.String(), b2.String())
    }

    if b1.Intersects(b3) {
        t.Errorf("%s intersects %s\n", b1.String(), b3.String())
    }
}
