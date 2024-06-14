package main

import (
	"fmt"
)

func old_interleave[E any](s []E, e E) []E {
	result := make([]E, 0, 2*len(s)-1)
	for i := range s {
		result = append(result, s[i])
		if i < len(s)-1 {
			result = append(result, e)
		}
	}
	return result
}

func interleave[E any](s []E, e E) func(func(int, E) bool) {
	return func(yield func(int, E) bool) {
		for i := range s {
			if !yield(2*i, s[i]) {
				return
			}
			if i < len(s)-1 {
				if !yield(2*i+1, e) {
					return
				}
			}
		}
	}
}

func main() {
	s := []string{"hello", "world", "this", "is", "a", "test"}
	for _, w := range interleave(s, " ") {
		fmt.Print(w)
	}

	// output: 'hello world this is a test'

	// old interleave (same effect but not lazy)
	// for _, w := range old_interleave(s, " ") {
	// 	fmt.Print(w)
	// }
}
