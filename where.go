// Package where  provides methods  for obtaining subsets from a given set.
package where

// A type, typically a collection, that satisfies subset.Interface
// is simply a collection of data indexed by the integers.
type Interface interface {
	Len() int
}

// condition reports whether the ith element satisfies an expression
type condition func(i int) bool

// Where returns the indices of the elements that satisfy a condition.
func Where(data Interface, c condition) []int {
	indicies := []int{}
	n := data.Len()
	for i := 0; i < n; i++ {
		// if ith element satisfies the condition, record.
		if c(i) {
			indicies = append(indicies, i)
		}
	}

	return indicies
}
