// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build cgo

package runtime

import _ "unsafe"

// used by cgo
//go:linkname _cgo_panic_internal
//go:linkname cgoAlwaysFalse
//go:linkname cgoUse
//go:linkname cgoKeepAlive
//go:linkname cgoCheckPointer
//go:linkname cgoCheckResult
//go:linkname cgoNoCallback
//go:linkname gobytes
//go:linkname gostringn
