package main

import (
	
	"testing"
)

// Тест для функции getCompareString с флагом -f
func TestGetCompareString_Fields(t *testing.T) {
	tests := []struct {
		name       string
		line       string
		numFields  int
		numChars   int
		ignoreCase bool
		want       string
	}{
		{
			name:       "skip 1 field",
			line:       "10 Alice Smith",
			numFields:  1,
			numChars:   0,
			ignoreCase: false,
			want:       "Alice Smith",
		},
		{
			name:       "skip 2 fields",
			line:       "10 20 Alice Smith",
			numFields:  2,
			numChars:   0,
			ignoreCase: false,
			want:       "Alice Smith",
		},
		{
			name:       "skip more fields than exist",
			line:       "A B",
			numFields:  5,
			numChars:   0,
			ignoreCase: false,
			want:       "",
		},
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

// Тест для флага -s (пропуск символов)
func TestGetCompareString_Chars(t *testing.T) {
	tests := []struct {
		name       string
		line       string
		numFields  int
		numChars   int
		ignoreCase bool
		want       string
	}{
		{
			name:       "skip 2 chars",
			line:       "Hello world",
			numFields:  0,
			numChars:   2,
			ignoreCase: false,
			want:       "llo world",
		},
		{
			name:       "skip more chars than line length",
			line:       "Hi",
			numFields:  0,
			numChars:   5,
			ignoreCase: false,
			want:       "",
		},
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

// Тест для флага -i (регистр)
func TestGetCompareString_Case(t *testing.T) {
	result := getCompareString("Hello", 0, 0, true)
	if result != "hello" {
		t.Errorf("getCompareString() = %q, want %q", result, "hello")
	}
}

// Тест: -f и -s вместе
func TestGetCompareString_FieldsAndChars(t *testing.T) {
	// "10 Alice Smith" → после -f 1: "Alice Smith" → после -s 2: "ice Smith"
	result := getCompareString("10 Alice Smith", 1, 2, false)
	if result != "ice Smith" {
		t.Errorf("getCompareString() = %q, want %q", result, "ice Smith")
	}
}

// Тест: -i и -f вместе
func TestGetCompareString_CaseAndFields(t *testing.T) {
	result := getCompareString("10 HELLO WORLD", 1, 0, true)
	if result != "hello world" {
		t.Errorf("getCompareString() = %q, want %q", result, "hello world")
	}
}