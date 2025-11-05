package main

import (
	"testing"
)

func TestGetCompareString_Flags(t *testing.T) {
	tests := []struct {
		name       string
		line       string
		numFields  int
		numChars   int
		ignoreCase bool
		want       string
	}{
		{"basic", "hello", 0, 0, false, "hello"},
		{"-f 1", "10 Alice", 1, 0, false, "Alice"},
		{"-f 2", "A B C D", 2, 0, false, "C D"},
		{"-s 2", "Hello", 0, 2, false, "llo"},
		{"-i", "HELLO", 0, 0, true, "hello"},
		{"-f 1 -s 1", "10 Alice", 1, 1, false, "lice"},
		{"-f too big", "A", 5, 0, false, ""},
		{"-s too big", "Hi", 0, 10, false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getCompareString(tt.line, tt.numFields, tt.numChars, tt.ignoreCase)
			if got != tt.want {
				t.Errorf("getCompareString() = %q, want %q", got, tt.want)
			}
		})
	}
}