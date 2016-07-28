package set

import (
	"bytes"
	"fmt"
)

// HashSet is a self-defined set data structure
type HashSet struct {
	m map[interface{}]bool
}

// NewHashSet initialize a new hash set
func NewHashSet() *HashSet {
	return &HashSet{m: make(map[interface{}]bool)}
}

// Add new element
func (set *HashSet) Add(e interface{}) bool {
	if !set.m[e] {
		set.m[e] = true
		return true
	}
	return false
}

// Remove eletement
func (set *HashSet) Remove(e interface{}) {
	delete(set.m, e)
}

// Clear all element by re-init a map
func (set *HashSet) Clear() {
	set.m = make(map[interface{}]bool)
}

// Contains test element in the set
func (set *HashSet) Contains(e interface{}) bool {
	return set.m[e]
}

// Len return the len of set
func (set *HashSet) Len() int {
	return len(set.m)
}

// Same check all keys in set are also in given set
func (set *HashSet) Same(other *HashSet) bool {
	if other == nil || set.Len() != other.Len() {
		return false
	}

	for key := range set.m {
		if !other.Contains(key) {
			return false
		}
	}

	return true
}

// Elements return all elements
func (set *HashSet) Elements() []interface{} {
	initialLen := len(set.m)
	snapshot := make([]interface{}, initialLen)
	actualLen := 0
	for key := range set.m {
		if actualLen < initialLen {
			snapshot[actualLen] = key
		} else {
			snapshot = append(snapshot, key)
		}
		actualLen++
	}

	if actualLen < initialLen {
		snapshot = snapshot[:actualLen]
	}
	return snapshot
}

// String content
func (set *HashSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("Set[")
	first := true
	for key := range set.m {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", key))
	}
	buf.WriteString("]")
	return buf.String()
}

// IsSuperset check is super set of other
func (set *HashSet) IsSuperset(other *HashSet) bool {
	if other == nil {
		return false
	}
	oneLen, otherLen := set.Len(), other.Len()
	if oneLen == 0 || oneLen == otherLen {
		return false
	}

	if oneLen > 0 && otherLen == 0 {
		return true
	}
	for _, v := range other.Elements() {
		if !set.Contains(v) {
			return false

		}
	}
	return true
}
