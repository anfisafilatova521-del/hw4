package main

import "testing"

func TestGetCompareString(t *testing.T) {
    tests := []struct {
        line       string
        f, s       int
        ignoreCase bool
        want       string
    }{
        {"10 Alice", 1, 0, false, "Alice"},
        {"HELLO", 0, 0, true, "hello"},
        {"A B C", 1, 1, false, " C"},
        {"Hi", 0, 5, false, ""},
    }
    for _, tt := range tests {
        got := getCompareString(tt.line, tt.f, tt.s, tt.ignoreCase)
        if got != tt.want {
            t.Errorf("getCompareString(%q, %d, %d, %t) = %q, want %q", tt.line, tt.f, tt.s, tt.ignoreCase, got, tt.want)
        }
    }
}