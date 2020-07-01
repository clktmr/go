// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package riscv64

import (
	"cmd/internal/obj/riscv"
	"cmd/internal/objabi"
	"cmd/internal/sys"
	"cmd/link/internal/ld"
	"cmd/link/internal/sym"
	"fmt"
	"log"
)

func lookupFuncSym(syms *sym.Symbols, name string) *sym.Symbol {
	if s := syms.ROLookup(name, sym.SymVerABI0); s != nil && s.FuncInfo != nil {
		return s
	}
	if s := syms.ROLookup(name, sym.SymVerABIInternal); s != nil && s.FuncInfo != nil {
		return s
	}
	return nil
}

func gentext(ctxt *ld.Link) {
	if ctxt.HeadType != objabi.Hnoos {
		return
	}
	vectors := ctxt.Syms.Lookup("runtime.vectors", 0)
	vectors.Type = sym.STEXT
	vectors.Attr |= sym.AttrReachable
	vectors.Align = 8

	unhandledInterrupt := ctxt.Syms.Lookup("runtime.unhandledExternalInterrupt", 0)
	if unhandledInterrupt == nil {
		ld.Errorf(nil, "runtime.unhandledExternalInterrupt not defined")
	}

	// search for user defined ISRs: //go:linkname functionName IRQ%d_Handler
	var irqHandlers [1024]*sym.Symbol // BUG: 1024 is PLIC specific
	irqNum := 1
	for i := 1; i < len(irqHandlers); i++ {
		s := lookupFuncSym(ctxt.Syms, ld.InterruptHandler(i))
		if s == nil {
			irqHandlers[i] = unhandledInterrupt
		} else {
			irqHandlers[i] = s
			irqNum = i + 1
		}
	}
	vectors.AddUint64(ctxt.Arch, uint64(irqNum))
	for _, s := range irqHandlers[1:irqNum] {
		s.Attr |= sym.AttrReachable
		rel := vectors.AddRel()
		rel.Off = int32(len(vectors.P))
		rel.Siz = 8
		rel.Type = objabi.R_ADDR
		rel.Sym = s
		vectors.AddUint64(ctxt.Arch, 0)
	}
	ctxt.Textp = append(ctxt.Textp, vectors)

	// move entry symbol on the beggining of text segment
	entry := ctxt.Syms.ROLookup(*ld.FlagEntrySymbol, sym.SymVerABI0)
	if entry == nil || entry.FuncInfo == nil {
		entry = ctxt.Syms.ROLookup(*ld.FlagEntrySymbol, sym.SymVerABIInternal)
		if entry == nil || entry.FuncInfo == nil {
			ld.Errorf(nil, "cannot find entry function: %s", *ld.FlagEntrySymbol)
		}
	}
	for i, s := range ctxt.Textp {
		if s == entry {
			copy(ctxt.Textp[1:], ctxt.Textp[:i])
			ctxt.Textp[0] = s
			return
		}
	}
	ld.Errorf(entry, "cannot find symbol in ctxt.Textp")
}

func adddynrela(ctxt *ld.Link, rel *sym.Symbol, s *sym.Symbol, r *sym.Reloc) {
	log.Fatalf("adddynrela not implemented")
}

func adddynrel(ctxt *ld.Link, s *sym.Symbol, r *sym.Reloc) bool {
	log.Fatalf("adddynrel not implemented")
	return false
}

func elfreloc1(ctxt *ld.Link, r *sym.Reloc, sectoff int64) bool {
	log.Fatalf("elfreloc1")
	return false
}

func elfsetupplt(ctxt *ld.Link) {
	log.Fatalf("elfsetuplt")
}

func machoreloc1(arch *sys.Arch, out *ld.OutBuf, s *sym.Symbol, r *sym.Reloc, sectoff int64) bool {
	log.Fatalf("machoreloc1 not implemented")
	return false
}

func archreloc(ctxt *ld.Link, r *sym.Reloc, s *sym.Symbol, val int64) (int64, bool) {
	switch r.Type {
	case objabi.R_CALLRISCV:
		// Nothing to do.
		return val, true

	case objabi.R_RISCV_PCREL_ITYPE, objabi.R_RISCV_PCREL_STYPE:
		pc := s.Value + int64(r.Off)
		off := ld.Symaddr(r.Sym) + r.Add - pc

		// Generate AUIPC and second instruction immediates.
		low, high, err := riscv.Split32BitImmediate(off)
		if err != nil {
			ld.Errorf(s, "R_RISCV_PCREL_ relocation does not fit in 32-bits: %d", off)
		}

		auipcImm, err := riscv.EncodeUImmediate(high)
		if err != nil {
			ld.Errorf(s, "cannot encode R_RISCV_PCREL_ AUIPC relocation offset for %s: %v", r.Sym.Name, err)
		}

		var secondImm, secondImmMask int64
		switch r.Type {
		case objabi.R_RISCV_PCREL_ITYPE:
			secondImmMask = riscv.ITypeImmMask
			secondImm, err = riscv.EncodeIImmediate(low)
			if err != nil {
				ld.Errorf(s, "cannot encode R_RISCV_PCREL_ITYPE I-type instruction relocation offset for %s: %v", r.Sym.Name, err)
			}
		case objabi.R_RISCV_PCREL_STYPE:
			secondImmMask = riscv.STypeImmMask
			secondImm, err = riscv.EncodeSImmediate(low)
			if err != nil {
				ld.Errorf(s, "cannot encode R_RISCV_PCREL_STYPE S-type instruction relocation offset for %s: %v", r.Sym.Name, err)
			}
		default:
			panic(fmt.Sprintf("Unknown relocation type: %v", r.Type))
		}

		auipc := int64(uint32(val))
		second := int64(uint32(val >> 32))

		auipc = (auipc &^ riscv.UTypeImmMask) | int64(uint32(auipcImm))
		second = (second &^ secondImmMask) | int64(uint32(secondImm))

		return second<<32 | auipc, true
	}

	return val, false
}

func archrelocvariant(ctxt *ld.Link, r *sym.Reloc, s *sym.Symbol, t int64) int64 {
	log.Fatalf("archrelocvariant")
	return -1
}

func asmb(ctxt *ld.Link) {
	if ctxt.IsELF {
		ld.Asmbelfsetup()
	}

	sect := ld.Segtext.Sections[0]
	ctxt.Out.SeekSet(int64(sect.Vaddr - ld.Segtext.Vaddr + ld.Segtext.Fileoff))
	ld.Codeblk(ctxt, int64(sect.Vaddr), int64(sect.Length))
	for _, sect = range ld.Segtext.Sections[1:] {
		ctxt.Out.SeekSet(int64(sect.Vaddr - ld.Segtext.Vaddr + ld.Segtext.Fileoff))
		ld.Datblk(ctxt, int64(sect.Vaddr), int64(sect.Length))
	}

	if ld.Segrodata.Filelen > 0 {
		ctxt.Out.SeekSet(int64(ld.Segrodata.Fileoff))
		ld.Datblk(ctxt, int64(ld.Segrodata.Vaddr), int64(ld.Segrodata.Filelen))
	}
	if ld.Segrelrodata.Filelen > 0 {
		ctxt.Out.SeekSet(int64(ld.Segrelrodata.Fileoff))
		ld.Datblk(ctxt, int64(ld.Segrelrodata.Vaddr), int64(ld.Segrelrodata.Filelen))
	}

	ctxt.Out.SeekSet(int64(ld.Segdata.Fileoff))
	ld.Datblk(ctxt, int64(ld.Segdata.Vaddr), int64(ld.Segdata.Filelen))

	ctxt.Out.SeekSet(int64(ld.Segdwarf.Fileoff))
	ld.Dwarfblk(ctxt, int64(ld.Segdwarf.Vaddr), int64(ld.Segdwarf.Filelen))
}

func asmb2(ctxt *ld.Link) {
	ld.Symsize = 0
	ld.Lcsize = 0
	symo := uint32(0)

	if !*ld.FlagS {
		if !ctxt.IsELF {
			ld.Errorf(nil, "unsupported executable format")
		}

		symo = uint32(ld.Segdwarf.Fileoff + ld.Segdwarf.Filelen)
		symo = uint32(ld.Rnd(int64(symo), int64(*ld.FlagRound)))
		ctxt.Out.SeekSet(int64(symo))

		ld.Asmelfsym(ctxt)
		ctxt.Out.Flush()
		ctxt.Out.Write(ld.Elfstrdat)

		if ctxt.LinkMode == ld.LinkExternal {
			ld.Elfemitreloc(ctxt)
		}
	}

	ctxt.Out.SeekSet(0)
	switch ctxt.HeadType {
	case objabi.Hlinux, objabi.Hnoos:
		ld.Asmbelf(ctxt, int64(symo))
	default:
		ld.Errorf(nil, "unsupported operating system")
	}
	ctxt.Out.Flush()

	if *ld.FlagC {
		fmt.Printf("textsize=%d\n", ld.Segtext.Filelen)
		fmt.Printf("datsize=%d\n", ld.Segdata.Filelen)
		fmt.Printf("bsssize=%d\n", ld.Segdata.Length-ld.Segdata.Filelen)
		fmt.Printf("symsize=%d\n", ld.Symsize)
		fmt.Printf("lcsize=%d\n", ld.Lcsize)
		fmt.Printf("total=%d\n", ld.Segtext.Filelen+ld.Segdata.Length+uint64(ld.Symsize)+uint64(ld.Lcsize))
	}
}
