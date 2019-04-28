package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// TestAdd test slice addition
func TestAdd(t *testing.T) {
	if diff := cmp.Diff([]int{7, 0, 8}, add([]int{2, 4, 3}, []int{5, 6, 4})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}

	if diff := cmp.Diff([]int{7, 4, 3}, add([]int{2, 4, 3}, []int{5})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}

	if diff := cmp.Diff([]int{9, 2, 6, 0, 0, 7, 1, 6, 2, 2, 4, 5, 6, 7, 8}, add([]int{3, 2, 2, 2, 5, 0, 4, 8, 6}, []int{6, 0, 4, 8, 4, 6, 7, 7, 5, 1, 4, 5, 6, 7, 8})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}
