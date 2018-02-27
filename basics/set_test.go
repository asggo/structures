package set

import (
	"sort"
	"testing"
)

func TestSet(t *testing.T) {
	s := NewSet([]int{})

	s.Add(1)
	if !s.Contains(1) {
		t.Error("1 is not a member of the set.")
	}

	s.Add(2)
	if !s.Contains(2) {
		t.Error("2 is not a member of the set.")
	}

	if s.Size() != 2 {
		t.Error("Set should contain two values.")
	}

	s.Remove(0)
	if s.Size() != 2 {
		t.Error("Set should contain two values.")
	}

	s.Remove(1)
	if s.Size() != 1 {
		t.Error("Set should contain one value.")
	}

	all := NewSet([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	members := all.Members()

	if !equalIntSlice(members, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Error("Set of all values does not contain the expected members.")
	}

	even := NewSet([]int{0, 2, 4, 6, 8})
	odd := NewSet([]int{1, 3, 5, 7, 9})
	intersect := Intersection(all, even)
	union := Union(all, odd)
	diff := Difference(all, even)

	if !Equal(intersect, even) {
		t.Error("The intersection of all and even should be even.")

	}

	if !Equal(union, all) {
		t.Error("The union of all and odd should be all.")
	}

	if !Equal(diff, odd) {
		t.Error("The difference of all and even should be odd.")
	}

}

func equalIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Ints(a)
	sort.Ints(b)

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
