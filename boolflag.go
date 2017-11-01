package veil

// BoolFlag is a flag that can either be a Boolean value or an int8.
type BoolFlag bool

// ToInt8 converts the flag to an int8.
func (b BoolFlag) ToInt8() int8 {
	if b {
		return 1
	}
	return 0
}

// ToBool converts the flag to a bool.
func (b BoolFlag) ToBool() bool {
	return bool(b)
}

// ToString converts the flag to a string representation.
func (b BoolFlag) ToString() string {
	if b {
		return "1"
	}
	return "0"
}
