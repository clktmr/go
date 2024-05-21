// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TODO: Do not use GOTARGET tags in runtime. Come up with something smarter
// based on some kind of memory description similar to the -M linker option.

//go:build noos && (imxrt1050 || imxrt1060)

package runtime

const (
	_OS                             = 0
	noos                            = true
	noosScaleDown                   = 4 // must be power of 2
	noosStackCacheSize              = 8 * 1024
	noosNumStackOrders              = 2
	noosHeapAddrBits                = 20         // enough for 1M OCRAM
	noosLogHeapArenaBytes           = 15         // 32 KiB
	noosArenaBaseOffset             = 0x20200000 // the begginning of OCRAM
	noosMinPhysPageSize             = 256
	noosSpanSetInitSpineCap         = 8
	noosStackMin                    = 1024
	noosStackSystem                 = 27 * 4 // register stacking at exception entry
	noosStackGuard                  = 464
	noosFinBlockSize                = 256
	noosSweepMinHeapDistance        = 1024
	noosDefaultHeapMinimum          = 8 * 1024
	noosMemoryLimitHeapGoalHeadroom = 1 << 16
	noosGCSweepBlockEntries         = 64
	noosGCSweepBufInitSpineCap      = 32
	noosGCBitsChunkBytes            = 2 * 1024
	noosSemTabSize                  = 31
	noosGOGC                        = 30
	noosTimeHistMaxBucketBits       = 46
)
