package randutils

import (
	_ "unsafe" // for runtime.fastrandn
)

// Uint32n returns a pseudorandom uint32 in [0,n).
//
//go:noescape
//go:linkname Uint32n runtime.fastrandn
func Uint32n(n uint32) uint32

// Uint32 returns a pseudorandom uint32.
//
//go:noescape
//go:linkname Uint32 runtime.fastrand
func Uint32() uint32
