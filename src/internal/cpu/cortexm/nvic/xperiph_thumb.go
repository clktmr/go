// DO NOT EDIT THIS FILE. GENERATED BY xgen.

package nvic

import (
	"embedded/mmio"
	"unsafe"
)

type Periph struct {
	ISER [16]RISER
	_    [16]uint32
	ICER [16]RICER
	_    [16]uint32
	ISPR [16]RISPR
	_    [16]uint32
	ICPR [16]RICPR
	_    [16]uint32
	IABR [16]RIABR
	_    [16]uint32
	ITNS [16]RITNS
	_    [16]uint32
	IPR  [496]RIPR
	_    [580]uint32
	STIR RSTIR
}

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

func NVIC() *Periph { return (*Periph)(unsafe.Pointer(uintptr(0xE000E100))) }

type ISER uint32

type RISER struct{ mmio.U32 }

func (r *RISER) Bits(mask ISER) ISER    { return ISER(r.U32.Bits(uint32(mask))) }
func (r *RISER) StoreBits(mask, b ISER) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RISER) SetBits(mask ISER)      { r.U32.SetBits(uint32(mask)) }
func (r *RISER) ClearBits(mask ISER)    { r.U32.ClearBits(uint32(mask)) }
func (r *RISER) Load() ISER             { return ISER(r.U32.Load()) }
func (r *RISER) Store(b ISER)           { r.U32.Store(uint32(b)) }

type RMISER struct{ mmio.UM32 }

func (rm RMISER) Load() ISER   { return ISER(rm.UM32.Load()) }
func (rm RMISER) Store(b ISER) { rm.UM32.Store(uint32(b)) }

type ICER uint32

type RICER struct{ mmio.U32 }

func (r *RICER) Bits(mask ICER) ICER    { return ICER(r.U32.Bits(uint32(mask))) }
func (r *RICER) StoreBits(mask, b ICER) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RICER) SetBits(mask ICER)      { r.U32.SetBits(uint32(mask)) }
func (r *RICER) ClearBits(mask ICER)    { r.U32.ClearBits(uint32(mask)) }
func (r *RICER) Load() ICER             { return ICER(r.U32.Load()) }
func (r *RICER) Store(b ICER)           { r.U32.Store(uint32(b)) }

type RMICER struct{ mmio.UM32 }

func (rm RMICER) Load() ICER   { return ICER(rm.UM32.Load()) }
func (rm RMICER) Store(b ICER) { rm.UM32.Store(uint32(b)) }

type ISPR uint32

type RISPR struct{ mmio.U32 }

func (r *RISPR) Bits(mask ISPR) ISPR    { return ISPR(r.U32.Bits(uint32(mask))) }
func (r *RISPR) StoreBits(mask, b ISPR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RISPR) SetBits(mask ISPR)      { r.U32.SetBits(uint32(mask)) }
func (r *RISPR) ClearBits(mask ISPR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RISPR) Load() ISPR             { return ISPR(r.U32.Load()) }
func (r *RISPR) Store(b ISPR)           { r.U32.Store(uint32(b)) }

type RMISPR struct{ mmio.UM32 }

func (rm RMISPR) Load() ISPR   { return ISPR(rm.UM32.Load()) }
func (rm RMISPR) Store(b ISPR) { rm.UM32.Store(uint32(b)) }

type ICPR uint32

type RICPR struct{ mmio.U32 }

func (r *RICPR) Bits(mask ICPR) ICPR    { return ICPR(r.U32.Bits(uint32(mask))) }
func (r *RICPR) StoreBits(mask, b ICPR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RICPR) SetBits(mask ICPR)      { r.U32.SetBits(uint32(mask)) }
func (r *RICPR) ClearBits(mask ICPR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RICPR) Load() ICPR             { return ICPR(r.U32.Load()) }
func (r *RICPR) Store(b ICPR)           { r.U32.Store(uint32(b)) }

type RMICPR struct{ mmio.UM32 }

func (rm RMICPR) Load() ICPR   { return ICPR(rm.UM32.Load()) }
func (rm RMICPR) Store(b ICPR) { rm.UM32.Store(uint32(b)) }

type IABR uint32

type RIABR struct{ mmio.U32 }

func (r *RIABR) Bits(mask IABR) IABR    { return IABR(r.U32.Bits(uint32(mask))) }
func (r *RIABR) StoreBits(mask, b IABR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RIABR) SetBits(mask IABR)      { r.U32.SetBits(uint32(mask)) }
func (r *RIABR) ClearBits(mask IABR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RIABR) Load() IABR             { return IABR(r.U32.Load()) }
func (r *RIABR) Store(b IABR)           { r.U32.Store(uint32(b)) }

type RMIABR struct{ mmio.UM32 }

func (rm RMIABR) Load() IABR   { return IABR(rm.UM32.Load()) }
func (rm RMIABR) Store(b IABR) { rm.UM32.Store(uint32(b)) }

type ITNS uint32

type RITNS struct{ mmio.U32 }

func (r *RITNS) Bits(mask ITNS) ITNS    { return ITNS(r.U32.Bits(uint32(mask))) }
func (r *RITNS) StoreBits(mask, b ITNS) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RITNS) SetBits(mask ITNS)      { r.U32.SetBits(uint32(mask)) }
func (r *RITNS) ClearBits(mask ITNS)    { r.U32.ClearBits(uint32(mask)) }
func (r *RITNS) Load() ITNS             { return ITNS(r.U32.Load()) }
func (r *RITNS) Store(b ITNS)           { r.U32.Store(uint32(b)) }

type RMITNS struct{ mmio.UM32 }

func (rm RMITNS) Load() ITNS   { return ITNS(rm.UM32.Load()) }
func (rm RMITNS) Store(b ITNS) { rm.UM32.Store(uint32(b)) }

type IPR uint8

type RIPR struct{ mmio.U8 }

func (r *RIPR) Bits(mask IPR) IPR     { return IPR(r.U8.Bits(uint8(mask))) }
func (r *RIPR) StoreBits(mask, b IPR) { r.U8.StoreBits(uint8(mask), uint8(b)) }
func (r *RIPR) SetBits(mask IPR)      { r.U8.SetBits(uint8(mask)) }
func (r *RIPR) ClearBits(mask IPR)    { r.U8.ClearBits(uint8(mask)) }
func (r *RIPR) Load() IPR             { return IPR(r.U8.Load()) }
func (r *RIPR) Store(b IPR)           { r.U8.Store(uint8(b)) }

type RMIPR struct{ mmio.UM8 }

func (rm RMIPR) Load() IPR   { return IPR(rm.UM8.Load()) }
func (rm RMIPR) Store(b IPR) { rm.UM8.Store(uint8(b)) }

type STIR uint32

type RSTIR struct{ mmio.U32 }

func (r *RSTIR) Bits(mask STIR) STIR    { return STIR(r.U32.Bits(uint32(mask))) }
func (r *RSTIR) StoreBits(mask, b STIR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSTIR) SetBits(mask STIR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSTIR) ClearBits(mask STIR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSTIR) Load() STIR             { return STIR(r.U32.Load()) }
func (r *RSTIR) Store(b STIR)           { r.U32.Store(uint32(b)) }

type RMSTIR struct{ mmio.UM32 }

func (rm RMSTIR) Load() STIR   { return STIR(rm.UM32.Load()) }
func (rm RMSTIR) Store(b STIR) { rm.UM32.Store(uint32(b)) }
