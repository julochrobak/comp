// Copyright (c) 2013 Julius Chrobak. You can use this source code
// under the terms of the MIT License found in the LICENSE file.

package main

import (
	. "math"
)

func min(a, b, c int) int {
	m := Min(float64(a), float64(b))
	return int(Min(m, float64(c)))
}

func dist(left, right string) int {
	if left == right {
		return 0
	}
	if len(left) == 0 {
		return len(right)
	}
	if len(right) == 0 {
		return len(left)
	}

	s := []rune(left)
	t := []rune(right)

	prev := make([]int, len(t)+1)
	curr := make([]int, len(t)+1)

	// initialize prev (simulating an empty s)
	for i := 0; i < len(prev); i++ {
		prev[i] = i
	}

	for i := 0; i < len(s); i++ {
		// first element of curr is A[i+1][0]
		//   edit distance is delete (i+1) chars from s to match empty t
		curr[0] = i + 1

		for j := 0; j < len(t); j++ {
			cost := 0
			if s[i] != t[j] {
				cost = 1
			}
			curr[j+1] = min(curr[j]+1, prev[j+1]+1, prev[j]+cost)
		}

		for j := 0; j < len(t); j++ {
			prev[j] = curr[j]
		}
	}

	return curr[len(t)]
}

func Fuzzy(left, right string) float64 {
	d := float64(dist(left, right))
	if d == 0 {
		return 1
	}

	s := []rune(left)
	t := []rune(right)
	l := Max(float64(len(s)), float64(len(t)))
	return (l - d) / l
}
