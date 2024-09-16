package utils

// StringSet struct definition
type StringSet map[string]struct{}

// NewStringSet function: creates and returns a new StringSet
func NewStringSet(vals []string) StringSet {
	set := make(StringSet)
	for _, val := range vals {
		set.Add(val)
	}
	return set
}

// Add method: add a new element to the StringSet
func (set StringSet) Add(s string) {
	set[s] = struct{}{}
}

// Remove method: remove an element from the StringSet
func (set StringSet) Remove(s string) {
	delete(set, s)
}

// Contains method: check if an element is in the StringSet
func (set StringSet) Contains(s string) bool {
	_, ok := set[s]
	return ok
}

// Difference method: returns the difference between two StringSets
func (set StringSet) Difference(other StringSet) StringSet {
	result := make(StringSet)
	for k := range set {
		if !other.Contains(k) {
			result.Add(k)
		}
	}
	return result
}

// Values method: returns the values of the StringSet
func (set StringSet) Values() []string {
	vals := make([]string, 0, len(set))
	for k := range set {
		vals = append(vals, k)
	}
	return vals
}
