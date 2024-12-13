// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtos

import (
	_ "unsafe"
)

//go:linkname notewakeup runtime.rtos_notewakeup
func notewakeup(n *Note)

//go:linkname notesleep runtime.rtos_notesleep
func notesleep(n *Note, timeout int64) bool

//go:linkname noteclear runtime.rtos_noteclear
func noteclear(n *Note) bool

func publicationBarrier()
