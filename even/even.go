// The even package implements a fast function for detecting if an integer
// is even or not.

package even

// Even returns true if i is even. Otherwise, false is returned
func Even(i int) bool {
	return i%2 == 0
}

// odd is the opposite of the Even
func odd(i int) bool {
	return i%2 == 1
}
