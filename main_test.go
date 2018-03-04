package main

import (
	"bytes"
	"testing"
)

func TestSkipToken(t *testing.T) {
	tests := []struct {
		in       []byte
		expected []byte
	}{
		{
			in:       []byte("'john' AND 1=1"),
			expected: []byte(" AND 1=1"),
		},
	}

	for _, test := range tests {
		got := SkipToken(test.in)
		if !bytes.Equal(got, test.expected) {
			t.Errorf("unexpected bytes. expected: %q, but got: %q", test.expected, got)
		}
	}
}
