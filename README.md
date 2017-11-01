# Veil


[![GoDoc](https://godoc.org/github.com/hscells/veil?status.svg)](https://godoc.org/github.com/hscells/veil)
[![Go Report Card](https://goreportcard.com/badge/github.com/hscells/veil)](https://goreportcard.com/report/github.com/hscells/veil)
[![cover.run](https://cover.run/go/github.com/hscells/veil.svg)](https://gocover.io/github.com/hscells/doi)


_A sensible and well tested flags and bitmask library for the go language._

Veil can be used anywhere the need for a bit mask (e.g. `1010`) needs to be used
to represent some state.

For instance, to find the values for a the mask above:

```go
fm := FlagsMapFromString([]string{"a", "b", "c", "d"}, "1010")
flag, err := fm.Get("a")
if err != nil {
    fmt.Println(flag)
}
```

A FlagsMap can also be constructed from a slice of `uint8`, a slice of `bool`,
or a map. Please see the [documentation](https://godoc.org/github.com/hscells/veil)
for the full reference.
