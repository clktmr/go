package mpu

import (
	"embedded/mmio"
	"unsafe"
)

type regs struct {
	typ  mmio.U32
	ctrl mmio.U32
	rnr  mmio.U32
	rbar mmio.U32
	rasr mmio.U32
}

func p() *regs { return (*regs)(unsafe.Pointer(uintptr(0xE000ED90))) }

// Type returns information about MPU unit:
// i - number of supported instruction regions,
// d - number of supported data regions.
// s - true if separate instruction and data regions are supported.
func Type() (i, d int, s bool) {
	typ := p().typ.Load()
	i = int(typ>>16) & 0xf
	d = int(typ>>8) & 0xf
	s = (typ&1 != 0)
	return
}

type Flags uint32

const (
	// If ENABLE is set MPU is enabled.
	ENABLE Flags = 1 << 0
	// If HFNMIENA is not set the MPU will be disabled during HardFault, NMI
	// and FAULTMASK handlers.
	HFNMIENA Flags = 1 << 1
	// If PRIVDEF is set the default memory map is used as background region for
	// privileged software access.
	PRIVDEFENA Flags = 1 << 2
)

// Set sets the flags specified by fl.
func Set(fl Flags) { p().ctrl.SetBits(uint32(fl)) }

// Clear clears the flags specified by fl.
func Clear(fl Flags) { p().ctrl.ClearBits(uint32(fl)) }

// State returns the current state.
func State() Flags { return Flags(p().ctrl.Load()) }

// Select selects the region number n.
func Select(n int) { p().rnr.Store(uint32(n)) }

type Attr uint32

const (
	ENA Attr = 1 << 0 // Enables region

	B Attr = 1 << 16 // Bufferable.
	C Attr = 1 << 17 // Cacheable.
	S Attr = 1 << 18 // Shareable.

	// Access permissons.
	Amask Attr = 7 << 24 // Use to extract access permission bits.
	A____ Attr = 0 << 24 // No access.
	Ar___ Attr = 5 << 24 // Priv-RO.
	Arw__ Attr = 1 << 24 // Priv-RW.
	Ar_r_ Attr = 6 << 24 // Priv-RO, Unpriv-RO.
	Arwr_ Attr = 2 << 24 // Priv-RW, Unpriv-RO.
	Arwrw Attr = 3 << 24 // Priv-RW, Unpriv-RW.

	XN Attr = 1 << 28 // Instruction access disable.
)

func SIZE(exp int) Attr {
	return Attr(exp-1) & 0x1f << 1
}

func (a Attr) SIZE() (exp int) {
	return int(a>>1)&0x1f + 1
}

func SRD(srd int) Attr {
	return Attr(srd&0xff) << 8
}

func (a Attr) SRD() int {
	return int(a>>8) & 0xff
}

func TEX(tex int) Attr {
	return Attr(tex&7) << 19
}

func (a Attr) TEX() int {
	return int(a>>19) & 7
}

const VALID = 1 << 4

func SetRegion(base uintptr, attr Attr) {
	p().rbar.Store(uint32(base))
	p().rasr.Store(uint32(attr))
}

func Region() (base uintptr, attr Attr) {
	return uintptr(p().rbar.Load()), Attr(p().rasr.Load())
}

/*
TODO: Implement SetRegions using STM instrucion.
type BaseAttr struct {
	RBAR uintptr
	RASR Attr
}

func SetRegions(bas [4]BaseAttr)
*/
