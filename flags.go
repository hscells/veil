package veil

// Flags is a slice containing our type.
type Flags []BoolFlag

// ToInt8 transforms the flags to a slice of bool.
func (flags Flags) ToInt8() (b []int8) {
	for _, flag := range flags {
		b = append(b, flag.ToInt8())
	}
	return
}

// ToBool transforms the flags to a slice of bool.
func (flags Flags) ToBool() (b []bool) {
	for _, flag := range flags {
		b = append(b, flag.ToBool())
	}
	return
}

// ToString creates the string representation of the flags.
func (flags Flags) ToString() (s string) {
	for _, flag := range flags {
		s += flag.ToString()
	}
	return
}

// NewFlags creates a new Flags using a list of bool.
func NewFlags(flags ...bool) (f Flags) {
	for _, flag := range flags {
		f = append(f, BoolFlag(flag))
	}
	return
}
