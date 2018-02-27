// The set package implements a set of integers.
package set

// Set defines a new set of integers.
type Set struct {
	members map[int]struct{}
}

// Add places a new member in the set.
func (s *Set) Add(i int) {
	_, ok := s.members[i]
	if !ok {
		s.members[i] = struct{}{}
	}
}

// Remove takes a member out of the set.
func (s *Set) Remove(i int) {
	delete(s.members, i)
}

// Contains returns true if i is a member of the set.
func (s *Set) Contains(i int) bool {
	_, ok := s.members[i]

	return ok
}

// Members returns a list of the members of the set.
func (s *Set) Members() []int {
	var members []int

	for i := range s.members {
		members = append(members, i)
	}

	return members
}

// Size returns the number of members in the set.
func (s *Set) Size() int {
	return len(s.members)
}

// Equal returns true if set x is the same length and contains all of the
// same members as the set y.
func Equal(x, y *Set) bool {
	if x.Size() != y.Size() {
		return false
	}

	for _, m := range x.Members() {
		if !y.Contains(m) {
			return false
		}
	}

	return true
}

// Union returns a new set that contains all of the members of x and all of
// The members of y.
func Union(x, y *Set) *Set {
	u := NewSet([]int{})

	for _, m := range x.Members() {
		u.Add(m)
	}
	for _, m := range y.Members() {
		u.Add(m)
	}

	return u
}

// Intersection returns a new set that contains only those members that appear
// in both x and y.
func Intersection(x, y *Set) *Set {
	i := NewSet([]int{})

	for _, m := range x.Members() {
		if y.Contains(m) {
			i.Add(m)
		}
	}

	return i
}

// Difference returns a new set that contains all of the members of x that do
// not appear in y.
func Difference(x, y *Set) *Set {
	d := NewSet([]int{})

	for _, m := range x.Members() {
		if !y.Contains(m) {
			d.Add(m)
		}
	}

	return d
}

// NewSet returns a new set of integers containing the integers in the members
// slice.
func NewSet(members []int) *Set {
	i := new(Set)
	i.members = make(map[int]struct{})

	for _, m := range members {
		i.Add(m)
	}

	return i
}
