// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import _ "unsafe"

// used in internal/godebug and syscall
//go:linkname write

// used in plugin
//go:linkname doInit

// used in math/bits
//go:linkname overflowError
//go:linkname divideError

// used in tests
//go:linkname extraMInUse
//go:linkname blockevent
//go:linkname haveHighResSleep
//go:linkname blockUntilEmptyFinalizerQueue
//go:linkname lockedOSThread
