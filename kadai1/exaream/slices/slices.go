// Package slices provides functions useful with slices.
package slices

// Check if a slice contains a string
func Contains(list []string, str string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
