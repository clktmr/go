// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package atomic

import (
	"internal/cpu"
	"internal/goos"
	"unsafe"
)

// Export some functions via linkname to assembly in sync/atomic.
//
//go:linkname Xchguintptr

type spinlock struct {
	v uint32
}

//go:nosplit
func (l *spinlock) lock() {
	for {
		if Cas(&l.v, 0, 1) {
			return
		}
	}
}

//go:nosplit
func (l *spinlock) unlock() {
	Store(&l.v, 0)
}

const _MCU = goos.IsNoos

var locktab [57*(1-_MCU) + 29*_MCU]struct {
	l   spinlock
	pad [cpu.CacheLinePadSize - unsafe.Sizeof(spinlock{})]byte
}

func addrLock(addr *uint64) *spinlock {
	return &locktab[(uintptr(unsafe.Pointer(addr))>>3)%uintptr(len(locktab))].l
}

const align64bit = 7

//go:nosplit
func goCas64(addr *uint64, old, new uint64) bool {
	if uintptr(unsafe.Pointer(addr))&align64bit != 0 {
		panicUnaligned()
	}
	_ = *addr // if nil, fault before taking the lock
	var ok bool
	addrLock(addr).lock()
	if *addr == old {
		*addr = new
		ok = true
	}
	addrLock(addr).unlock()
	return ok
}

//go:nosplit
func goXadd64(addr *uint64, delta int64) uint64 {
	if uintptr(unsafe.Pointer(addr))&align64bit != 0 {
		panicUnaligned()
	}
	_ = *addr // if nil, fault before taking the lock
	var r uint64
	addrLock(addr).lock()
	r = *addr + uint64(delta)
	*addr = r
	addrLock(addr).unlock()
	return r
}

//go:nosplit
func goXchg64(addr *uint64, v uint64) uint64 {
	if uintptr(unsafe.Pointer(addr))&align64bit != 0 {
		panicUnaligned()
	}
	_ = *addr // if nil, fault before taking the lock
	var r uint64
	addrLock(addr).lock()
	r = *addr
	*addr = v
	addrLock(addr).unlock()
	return r
}

//go:nosplit
func goLoad64(addr *uint64) uint64 {
	if uintptr(unsafe.Pointer(addr))&align64bit != 0 {
		panicUnaligned()
	}
	_ = *addr // if nil, fault before taking the lock
	var r uint64
	addrLock(addr).lock()
	r = *addr
	addrLock(addr).unlock()
	return r
}

//go:nosplit
func goStore64(addr *uint64, v uint64) {
	if uintptr(unsafe.Pointer(addr))&align64bit != 0 {
		panicUnaligned()
	}
	_ = *addr // if nil, fault before taking the lock
	addrLock(addr).lock()
	*addr = v
	addrLock(addr).unlock()
}

//go:nosplit
func Or8(addr *uint8, v uint8) {
	// Align down to 4 bytes and use 32-bit CAS.
	uaddr := uintptr(unsafe.Pointer(addr))
	addr32 := (*uint32)(unsafe.Pointer(uaddr &^ 3))
	word := uint32(v) << ((uaddr & 3) * 8) // little endian
	for {
		old := *addr32
		if Cas(addr32, old, old|word) {
			return
		}
	}
}

//go:nosplit
func And8(addr *uint8, v uint8) {
	// Align down to 4 bytes and use 32-bit CAS.
	uaddr := uintptr(unsafe.Pointer(addr))
	addr32 := (*uint32)(unsafe.Pointer(uaddr &^ 3))
	word := uint32(v) << ((uaddr & 3) * 8)    // little endian
	mask := uint32(0xFF) << ((uaddr & 3) * 8) // little endian
	word |= ^mask
	for {
		old := *addr32
		if Cas(addr32, old, old&word) {
			return
		}
	}
}

//go:nosplit
func Or(addr *uint32, v uint32)

//go:nosplit
func And(addr *uint32, v uint32)

//go:noescape
func Load(addr *uint32) uint32

// NO go:noescape annotation; *addr escapes if result escapes (#31525)
func Loadp(addr unsafe.Pointer) unsafe.Pointer

//go:noescape
func Load64(addr *uint64) uint64

//go:noescape
func Load8(addr *uint8) uint8

//go:noescape
func LoadAcq(addr *uint32) uint32

//go:noescape
func LoadAcquintptr(ptr *uintptr) uintptr

// Not noescape -- it installs a pointer to addr.
func StorepNoWB(addr unsafe.Pointer, v unsafe.Pointer)

//go:noescape
func Store(addr *uint32, v uint32)

//go:noescape
func StoreRel(addr *uint32, v uint32)

//go:noescape
func StoreReluintptr(addr *uintptr, v uintptr)

//go:noescape
func Store64(addr *uint64, v uint64)

//go:noescape
func Store8(addr *uint8, v uint8)

//go:noescape
func Cas64(addr *uint64, old, new uint64) bool

//go:noescape
func CasRel(addr *uint32, old, new uint32) bool

//go:noescape
func Xadd(val *uint32, delta int32) uint32

//go:noescape
func Xadd64(addr *uint64, delta int64) uint64

//go:noescape
func Xchg64(addr *uint64, v uint64) uint64

//go:noescape
func Xadduintptr(val *uintptr, delta uintptr) uintptr

//go:noescape
func Xchg(addr *uint32, v uint32) uint32

//go:nosplit
func Xchguintptr(addr *uintptr, v uintptr) uintptr {
	return uintptr(Xchg((*uint32)(unsafe.Pointer(addr)), uint32(v)))
}
