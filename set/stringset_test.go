package set

import (
	"testing"
)

func TestStringSet(t *testing.T) {
	s := NewStringSet([]string{})

	s.Add("1")
	if !s.Contains("1") {
		t.Error("1 is not a member of the set.")
	}

	s.Add("2")
	if !s.Contains("2") {
		t.Error("2 is not a member of the set.")
	}

	if s.Size() != 2 {
		t.Error("Set should contain two values.")
	}

	s.Remove("0")
	if s.Size() != 2 {
		t.Error("Set should contain two values.")
	}

	s.Remove("1")
	if s.Size() != 1 {
		t.Error("Set should contain one value.")
	}

	allSlice := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	all := NewStringSet(allSlice)
	members := all.Members()

	if !equalStringSlice(members, allSlice) {
		t.Error("Set of all values does not contain the expected members.")
	}

	even := NewStringSet([]string{"0", "2", "4", "6", "8"})
	odd := NewStringSet([]string{"1", "3", "5", "7", "9"})
	intersect := all.Intersection(even)
	union := all.Union(odd)
	diff := all.Difference(even)

	if !even.Subset(all) {
		t.Error("Even should be a subset of all.")
	}

	if all.Subset(odd) {
		t.Error("All should not be a subset of odd.")
	}

	if !intersect.Equal(even) {
		t.Error("The intersection of all and even should be even.")
	}

	if !union.Equal(all) {
		t.Error("The union of all and odd should be all.")
	}

	if !diff.Equal(odd) {
		t.Error("The difference of all and even should be odd.")
	}
}

func equalStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
