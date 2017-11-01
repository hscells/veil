// Package veil is a sensible flag and bit mask library.
package veil

import "github.com/pkg/errors"

// FlagsMap is a container for representing named flags.
type FlagsMap struct {
	Keys  []string
	Flags Flags
}

// Zip transforms the flags and strings into a native go map.
func (b FlagsMap) Zip() (m map[string]bool) {
	m = map[string]bool{}
	for index, key := range b.Keys {
		m[key] = b.Flags[index].ToBool()
	}
	return
}

// ToString is the string encoding of the map. In this representation the keys are lost.
func (b FlagsMap) ToString() string {
	return b.Flags.ToString()
}

// Set updates the values stored for an existing key.
func (b FlagsMap) Set(k string, flag bool) error {
	for index, key := range b.Keys {
		if key == k {
			b.Flags[index] = BoolFlag(flag)
			return nil
		}
	}
	return errors.New("key not found")
}

// Get retrieves a flag for a given key.
func (b FlagsMap) Get(k string) (bool, error) {
	for index, key := range b.Keys {
		if key == k {
			return b.Flags[index].ToBool(), nil
		}
	}
	return false, errors.New("key not found")
}

// NewFlagsMap creates a new FlagsMap using the supplied keys and a list of Boolean values. If there are less flags than
// keys, the flags will be added for the first keys and the rest default to 0 (for instance setting the first few keys
// to their "on" states.
func NewFlagsMap(keys []string, flags ...bool) FlagsMap {
	// We have an empty map.
	if len(keys) == 0 {
		return FlagsMap{}
	}
	// There is not an equal number of keys to flags, but we can create it partially.
	if len(keys) != len(flags) {
		bfm := FlagsMap{Keys: keys, Flags: []BoolFlag{}}
		// We add the flags that might be supplied.
		for index, flag := range flags {
			// And we leave early if there are too many flags.
			if index >= len(keys) {
				return bfm
			}
			bfm.Flags = append(bfm.Flags, BoolFlag(flag))
		}
		// We add zeros to the ones not added.
		for i := len(bfm.Flags) - 1; i < len(keys)-1; i++ {
			bfm.Flags = append(bfm.Flags, false)
		}
		return bfm
	}
	// Otherwise we can just make the new map.
	boolFlags := []BoolFlag{}
	for _, flag := range flags {
		boolFlags = append(boolFlags, BoolFlag(flag))
	}
	return FlagsMap{Keys: keys, Flags: boolFlags}
}

// DecodeFlagsMap creates a new FlagsMap using the supplied keys and a slice of flags (int8) in the same way as
// NewFlagsMap creates the object.
func DecodeFlagsMap(keys []string, flags []int8) FlagsMap {
	boolFlags := []bool{}
	for _, flag := range flags {
		var b bool
		if flag > 0 {
			b = true
		} else {
			b = false
		}
		boolFlags = append(boolFlags, b)
	}
	return NewFlagsMap(keys, boolFlags...)
}

// FlagsMapFromMap creates a FlagsMap from a map.
func FlagsMapFromMap(m map[string]bool) FlagsMap {
	keys := []string{}
	flags := []int8{}
	for k, v := range m {
		keys = append(keys, k)
		flags = append(flags, BoolFlag(v).ToInt8())
	}
	return DecodeFlagsMap(keys, flags)
}

// FlagsMapFromString creates a FlagsMap from a slice of keys and a string mask.
func FlagsMapFromString(keys []string, mask string) FlagsMap {
	flags := []int8{}
	for _, char := range mask {
		if char == '0' {
			flags = append(flags, 0)
		} else {
			flags = append(flags, 1)
		}
	}
	return DecodeFlagsMap(keys, flags)
}
