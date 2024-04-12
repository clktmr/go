// DO NOT EDIT THIS FILE. GENERATED BY xgen.

package acc

import (
	"embedded/mmio"
	"unsafe"
)

type Periph struct {
	ITCMCR RITCMCR
	DTCMCR RDTCMCR
	AHBPCR RAHBPCR
	CACR   RCACR
	AHBSCR RAHBSCR
	_      uint32
	ABFSR  RABFSR
}

func ACC() *Periph { return (*Periph)(unsafe.Pointer(uintptr(0xE000EF90))) }

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type ITCMCR uint32

type RITCMCR struct{ mmio.U32 }

func (r *RITCMCR) LoadBits(mask ITCMCR) ITCMCR { return ITCMCR(r.U32.LoadBits(uint32(mask))) }
func (r *RITCMCR) StoreBits(mask, b ITCMCR)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RITCMCR) SetBits(mask ITCMCR)         { r.U32.SetBits(uint32(mask)) }
func (r *RITCMCR) ClearBits(mask ITCMCR)       { r.U32.ClearBits(uint32(mask)) }
func (r *RITCMCR) Load() ITCMCR                { return ITCMCR(r.U32.Load()) }
func (r *RITCMCR) Store(b ITCMCR)              { r.U32.Store(uint32(b)) }

type RMITCMCR struct{ mmio.UM32 }

func (rm RMITCMCR) Load() ITCMCR   { return ITCMCR(rm.UM32.Load()) }
func (rm RMITCMCR) Store(b ITCMCR) { rm.UM32.Store(uint32(b)) }

func ITCMEN_(p *Periph) RMITCMCR {
	return RMITCMCR{mmio.UM32{&p.ITCMCR.U32, uint32(ITCMEN)}}
}

func ITCMRMW_(p *Periph) RMITCMCR {
	return RMITCMCR{mmio.UM32{&p.ITCMCR.U32, uint32(ITCMRMW)}}
}

func ITCMRETEN_(p *Periph) RMITCMCR {
	return RMITCMCR{mmio.UM32{&p.ITCMCR.U32, uint32(ITCMRETEN)}}
}

func ITCMSZ_(p *Periph) RMITCMCR {
	return RMITCMCR{mmio.UM32{&p.ITCMCR.U32, uint32(ITCMSZ)}}
}

type DTCMCR uint32

type RDTCMCR struct{ mmio.U32 }

func (r *RDTCMCR) LoadBits(mask DTCMCR) DTCMCR { return DTCMCR(r.U32.LoadBits(uint32(mask))) }
func (r *RDTCMCR) StoreBits(mask, b DTCMCR)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDTCMCR) SetBits(mask DTCMCR)         { r.U32.SetBits(uint32(mask)) }
func (r *RDTCMCR) ClearBits(mask DTCMCR)       { r.U32.ClearBits(uint32(mask)) }
func (r *RDTCMCR) Load() DTCMCR                { return DTCMCR(r.U32.Load()) }
func (r *RDTCMCR) Store(b DTCMCR)              { r.U32.Store(uint32(b)) }

type RMDTCMCR struct{ mmio.UM32 }

func (rm RMDTCMCR) Load() DTCMCR   { return DTCMCR(rm.UM32.Load()) }
func (rm RMDTCMCR) Store(b DTCMCR) { rm.UM32.Store(uint32(b)) }

func DTCMEN_(p *Periph) RMDTCMCR {
	return RMDTCMCR{mmio.UM32{&p.DTCMCR.U32, uint32(DTCMEN)}}
}

func DTCMRMW_(p *Periph) RMDTCMCR {
	return RMDTCMCR{mmio.UM32{&p.DTCMCR.U32, uint32(DTCMRMW)}}
}

func DTCMRETEN_(p *Periph) RMDTCMCR {
	return RMDTCMCR{mmio.UM32{&p.DTCMCR.U32, uint32(DTCMRETEN)}}
}

func DTCMSZ_(p *Periph) RMDTCMCR {
	return RMDTCMCR{mmio.UM32{&p.DTCMCR.U32, uint32(DTCMSZ)}}
}

type AHBPCR uint32

type RAHBPCR struct{ mmio.U32 }

func (r *RAHBPCR) LoadBits(mask AHBPCR) AHBPCR { return AHBPCR(r.U32.LoadBits(uint32(mask))) }
func (r *RAHBPCR) StoreBits(mask, b AHBPCR)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAHBPCR) SetBits(mask AHBPCR)         { r.U32.SetBits(uint32(mask)) }
func (r *RAHBPCR) ClearBits(mask AHBPCR)       { r.U32.ClearBits(uint32(mask)) }
func (r *RAHBPCR) Load() AHBPCR                { return AHBPCR(r.U32.Load()) }
func (r *RAHBPCR) Store(b AHBPCR)              { r.U32.Store(uint32(b)) }

type RMAHBPCR struct{ mmio.UM32 }

func (rm RMAHBPCR) Load() AHBPCR   { return AHBPCR(rm.UM32.Load()) }
func (rm RMAHBPCR) Store(b AHBPCR) { rm.UM32.Store(uint32(b)) }

func AHBPEN_(p *Periph) RMAHBPCR {
	return RMAHBPCR{mmio.UM32{&p.AHBPCR.U32, uint32(AHBPEN)}}
}

func AHBPSZ_(p *Periph) RMAHBPCR {
	return RMAHBPCR{mmio.UM32{&p.AHBPCR.U32, uint32(AHBPSZ)}}
}

type CACR uint32

type RCACR struct{ mmio.U32 }

func (r *RCACR) LoadBits(mask CACR) CACR { return CACR(r.U32.LoadBits(uint32(mask))) }
func (r *RCACR) StoreBits(mask, b CACR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCACR) SetBits(mask CACR)       { r.U32.SetBits(uint32(mask)) }
func (r *RCACR) ClearBits(mask CACR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RCACR) Load() CACR              { return CACR(r.U32.Load()) }
func (r *RCACR) Store(b CACR)            { r.U32.Store(uint32(b)) }

type RMCACR struct{ mmio.UM32 }

func (rm RMCACR) Load() CACR   { return CACR(rm.UM32.Load()) }
func (rm RMCACR) Store(b CACR) { rm.UM32.Store(uint32(b)) }

func SIWT_(p *Periph) RMCACR {
	return RMCACR{mmio.UM32{&p.CACR.U32, uint32(SIWT)}}
}

func ECCDIS_(p *Periph) RMCACR {
	return RMCACR{mmio.UM32{&p.CACR.U32, uint32(ECCDIS)}}
}

func FORCEWT_(p *Periph) RMCACR {
	return RMCACR{mmio.UM32{&p.CACR.U32, uint32(FORCEWT)}}
}

type AHBSCR uint32

type RAHBSCR struct{ mmio.U32 }

func (r *RAHBSCR) LoadBits(mask AHBSCR) AHBSCR { return AHBSCR(r.U32.LoadBits(uint32(mask))) }
func (r *RAHBSCR) StoreBits(mask, b AHBSCR)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAHBSCR) SetBits(mask AHBSCR)         { r.U32.SetBits(uint32(mask)) }
func (r *RAHBSCR) ClearBits(mask AHBSCR)       { r.U32.ClearBits(uint32(mask)) }
func (r *RAHBSCR) Load() AHBSCR                { return AHBSCR(r.U32.Load()) }
func (r *RAHBSCR) Store(b AHBSCR)              { r.U32.Store(uint32(b)) }

type RMAHBSCR struct{ mmio.UM32 }

func (rm RMAHBSCR) Load() AHBSCR   { return AHBSCR(rm.UM32.Load()) }
func (rm RMAHBSCR) Store(b AHBSCR) { rm.UM32.Store(uint32(b)) }

func CTL_(p *Periph) RMAHBSCR {
	return RMAHBSCR{mmio.UM32{&p.AHBSCR.U32, uint32(CTL)}}
}

func TPRI_(p *Periph) RMAHBSCR {
	return RMAHBSCR{mmio.UM32{&p.AHBSCR.U32, uint32(TPRI)}}
}

func INITCOUNT_(p *Periph) RMAHBSCR {
	return RMAHBSCR{mmio.UM32{&p.AHBSCR.U32, uint32(INITCOUNT)}}
}

type ABFSR uint32

type RABFSR struct{ mmio.U32 }

func (r *RABFSR) LoadBits(mask ABFSR) ABFSR { return ABFSR(r.U32.LoadBits(uint32(mask))) }
func (r *RABFSR) StoreBits(mask, b ABFSR)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RABFSR) SetBits(mask ABFSR)        { r.U32.SetBits(uint32(mask)) }
func (r *RABFSR) ClearBits(mask ABFSR)      { r.U32.ClearBits(uint32(mask)) }
func (r *RABFSR) Load() ABFSR               { return ABFSR(r.U32.Load()) }
func (r *RABFSR) Store(b ABFSR)             { r.U32.Store(uint32(b)) }

type RMABFSR struct{ mmio.UM32 }

func (rm RMABFSR) Load() ABFSR   { return ABFSR(rm.UM32.Load()) }
func (rm RMABFSR) Store(b ABFSR) { rm.UM32.Store(uint32(b)) }

func ITCM_(p *Periph) RMABFSR {
	return RMABFSR{mmio.UM32{&p.ABFSR.U32, uint32(ITCM)}}
}

func DTCM_(p *Periph) RMABFSR {
	return RMABFSR{mmio.UM32{&p.ABFSR.U32, uint32(DTCM)}}
}

func AHBP_(p *Periph) RMABFSR {
	return RMABFSR{mmio.UM32{&p.ABFSR.U32, uint32(AHBP)}}
}

func AXIM_(p *Periph) RMABFSR {
	return RMABFSR{mmio.UM32{&p.ABFSR.U32, uint32(AXIM)}}
}

func EPPB_(p *Periph) RMABFSR {
	return RMABFSR{mmio.UM32{&p.ABFSR.U32, uint32(EPPB)}}
}

func AXIMTYPE_(p *Periph) RMABFSR {
	return RMABFSR{mmio.UM32{&p.ABFSR.U32, uint32(AXIMTYPE)}}
}
