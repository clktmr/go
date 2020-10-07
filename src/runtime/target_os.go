// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !noos

package runtime

const (
	_OS                        = 1
	noos                       = false
	noosStackCacheSize         = 0
	noosNumStackOrders         = 0
	noosHeapAddrBits           = 0
	noosLogHeapArenaBytes      = 0
	noosArenaBaseOffset        = 0
	noosMinPhysPageSize        = 0
	noosFinBlockSize           = 0
	noosSweepMinHeapDistance   = 0
	noosDefaultHeapMinimum     = 0
	noosHeapGoalInc            = 0
	noosGCSweepBlockEntries    = 0
	noosGCSweepBufInitSpineCap = 0
	noosGCBitsChunkBytes       = 0
)
