// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build 386 amd64 amd64p32 s390x arm thumb arm64 ppc64 ppc64le mips mipsle wasm

package bytealg

import _ "unsafe" // For go:linkname

//go:noescape
func Compare(a, b []byte) int

// The declaration below generates ABI wrappers for functions
// implemented in assembly in this package but declared in another
// package.

//go:linkname abigen_runtime_cmpstring runtime.cmpstring
func abigen_runtime_cmpstring(a, b string) int
