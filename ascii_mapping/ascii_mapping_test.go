package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// TestMapping testing ascii mapping algo
func TestMapping(t *testing.T) {
	if diff := cmp.Diff(float64(100), mapping("Hello")); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}

	if diff := cmp.Diff(float64(104), mapping("Test")); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}

	if diff := cmp.Diff(float64(82), mapping("QWERTY")); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}

	if diff := cmp.Diff(float64(48.142857142857146), mapping("!@#$%^&")); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}

	if diff := cmp.Diff(float64(87.83333333333333), mapping("This is fun!")); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}
