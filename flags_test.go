package veil

import "testing"

func TestNewFlagsBool(t *testing.T) {
	f := NewFlags(true, false, true, false)

	got := f.ToBool()
	expected := []bool{true, false, true, false}
	for i, g := range got {
		if g != expected[i] {
			t.Errorf("expected %v, got %v", expected, got)
		}
	}
}

func TestNewFlagsInt8(t *testing.T) {
	f := NewFlags(true, false, true, false)

	got := f.ToInt8()
	expected := []int8{1, 0, 1, 0}
	for i, g := range got {
		if g != expected[i] {
			t.Errorf("expected %v, got %v", expected, got)
		}
	}
}

func TestNewFlagsString(t *testing.T) {
	f := NewFlags(true, false, true, false)

	got := f.ToString()
	expected := "1010"
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
