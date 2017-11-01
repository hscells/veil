package veil

import "testing"

var (
	keys = []string{"pencil", "eraser", "ruler", "staple"}
	mask = "1010"
)

func TestFlagsMapFromString_Success(t *testing.T) {
	fm := FlagsMapFromString(keys, mask)

	got := fm.ToString()
	expected := "1010"
	if got != expected {
		t.Errorf("expected %v, got %v", mask, got)
	}
}

func TestFlagsMapFromMap_Success(t *testing.T) {
	m := map[string]bool{
		"pencil": true,
		"eraser": false,
		"ruler":  true,
		"staple": false,
	}

	fm := FlagsMapFromMap(m)

	got, err := fm.Get("pencil")
	if err != nil {
		t.Fatal(err)
	}

	if got != true {
		t.Errorf("expected %v, got %v", true, got)
	}
}

func TestDecodeFlagsMap_Success(t *testing.T) {
	fm := DecodeFlagsMap(keys, []int8{1, 0, 1, 0})

	got := fm.ToString()
	expected := "1010"
	if got != expected {
		t.Errorf("expected %v, got %v", mask, got)
	}
}

func TestNewFlagsMap1_Success(t *testing.T) {
	fm := NewFlagsMap(keys, true, false, true, false)

	got := fm.ToString()
	expected := mask
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestNewFlagsMap2_Success(t *testing.T) {
	fm := NewFlagsMap(keys, true, true)

	got := fm.ToString()
	expected := "1100"
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestNewFlagsMap3_Success(t *testing.T) {
	fm := NewFlagsMap(keys)

	got := fm.ToString()
	expected := "0000"
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestNewFlagsMap4_Success(t *testing.T) {
	fm := NewFlagsMap([]string{})

	got := fm.ToString()
	expected := ""
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestNewFlagsMap5_Success(t *testing.T) {
	fm := NewFlagsMap(keys, true, true, true, true, true, true, true)

	got := fm.ToString()
	expected := "1111"
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestFlagsMap_Get1_Success(t *testing.T) {
	fm := FlagsMapFromString(keys, mask)

	got, err := fm.Get("pencil")
	if err != nil {
		t.Fatal(err)
	}

	if got != true {
		t.Errorf("expected %v, got %v", true, got)
	}
}

func TestFlagsMap_Get2_Success(t *testing.T) {
	fm := FlagsMapFromString(keys, mask)

	got, err := fm.Get("staple")
	if err != nil {
		t.Fatal(err)
	}

	if got != false {
		t.Errorf("expected %v, got %v", false, got)
	}
}

func TestFlagsMap_Get3_Failure(t *testing.T) {
	fm := FlagsMapFromString(keys, mask)

	_, err := fm.Get("garbage")
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestFlagsMap_Set1_Success(t *testing.T) {
	fm := FlagsMapFromString(keys, mask)

	err := fm.Set("staple", true)
	if err != nil {
		t.Fatal(err)
	}

	got, err := fm.Get("staple")
	if err != nil {
		t.Fatal(err)
	}

	if got != true {
		t.Errorf("expected %v, got %v", true, got)
	}
}

func TestFlagsMap_Set2_Failure(t *testing.T) {
	fm := FlagsMapFromString(keys, mask)

	err := fm.Set("garbage", true)
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestFlagsMap_Zip(t *testing.T) {
	fm := FlagsMapFromString(keys, mask)

	m := fm.Zip()

	got := m["pencil"]
	if got != true {
		t.Errorf("expected %v, got %v", true, got)
	}
}
