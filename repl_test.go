package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCleanInput(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected []string
	}{
		"simple": {
			input:    "hello there",
			expected: []string{"hello", "there"},
		},
		"toLower": {
			input:    "HELLo there",
			expected: []string{"hello", "there"},
		},
		"trailingWithespace": {
			input:    "hello there ",
			expected: []string{"hello", "there"},
		},
		"leadingWithespace": {
			input:    " hello there",
			expected: []string{"hello", "there"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := cleanInput(tc.input)
			diff := cmp.Diff(got, tc.expected)
			if diff != "" {
				t.Fatalf("%s", diff)
			}
		})
	}
}
