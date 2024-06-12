// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "go_tls.h"
#include "funcdata.h"
#include "textflag.h"
#include "asm_mips64.h"

// BREAK is overloaded CACHE opcode. Register number specifies the cache op.
#define CACHE BREAK

TEXT _rt0_mips64_noos(SB),NOSPLIT|NOFRAME,$0
	JMP ·rt0_target(SB)

TEXT runtime·_rt0_mips64_noos1(SB),NOSPLIT|NOFRAME,$0
	// Clear .bss, .noptrbss and unallocated memory.
	SUBU $16, R29

	MOVW $runtime·bss(SB), R4
	MOVW $runtime·ebss(SB), R5
	SUB  R4, R5
	MOVV R4, 8(R29)
	MOVV R5, 16(R29)
	JAL  runtime·memclrNoHeapPointers(SB)

	MOVW $runtime·noptrbss(SB), R4
	MOVW $runtime·enoptrbss(SB), R5
	SUB  R4, R5
	MOVV R4, 8(R29)
	MOVV R5, 16(R29)
	JAL  runtime·memclrNoHeapPointers(SB)

	MOVW $runtime·end(SB), R4
	MOVW $runtime·ramend(SB), R5
	SUB  R4, R5
	MOVV R4, 8(R29)
	MOVV R5, 16(R29)
	JAL  runtime·memclrNoHeapPointers(SB)

	ADDU $16, R29

	// Load interrupt vector
	MOVW $runtime·intvector(SB), R8
	MOVW $0xa0000000, R9
	MOVW $4, R10
loop:
	MOVW (R8), R11
	MOVW R11, 0(R9)
	MOVW R11, 0x80(R9)
	MOVW R11, 0x100(R9)
	MOVW R11, 0x180(R9)
	CACHE R16, 0(R9)
	CACHE R16, 0x80(R9)
	CACHE R16, 0x100(R9)
	CACHE R16, 0x180(R9)
	ADD $4, R8
	ADD $4, R9
	ADDU $-1, R10
	BGTZ R10,loop

	JMP runtime·rt0_go(SB)


#define PALLOC_MIN 20*1024

TEXT runtime·rt0_go(SB),NOSPLIT|NOFRAME|TOPFRAME,$0
	// setup main stack in cpu0.gh
	MOVV  $runtime·cpu0(SB), R8  // gh is the first field of the cpuctx struct
	MOVV  $runtime·ramend(SB), R29  // main stack starts at the end of memory
	SUB   $16, R29
	MOVV  R29, (g_stack+stack_hi)(R8)
	SUB   $4096, R29, R9
	MOVV  R9, (g_stack+stack_lo)(R8)
	ADD   $const_stackGuard, R9
	MOVV  R9, g_stackguard0(R8)
	MOVV  R9, g_stackguard1(R8)

	// set up m0 (bootstrap thread), temporarily use cpu0.gh as g
	MOVV  $runtime·m0(SB), R9
	MOVV  R8, m_g0(R9)  // m0.g0 = cpu0.gh
	MOVV  R9, g_m(R8)   // cpu0.gh.m = m0

	MOVV  R8, g  // set g to gh

	JAL runtime·check(SB)
	JAL runtime·osinit(SB)

	// initialize noosMem

	MOVV  $runtime·end(SB), R8
	MOVV  $runtime·ramend(SB), R9
	SUB   $4096, R9
	SUB   R8, R9, R13  // size of available memory (DMA capable)

	// estimate the space need for non-heap allocations
	SRL   $(const__PageShift), R13, R12
	MOVV  $mspan__size, R10
	MUL   R10, R12
	MOVV  LO, R12
	ADD   $PALLOC_MIN, R12

	MOVV  $runtime·nodmastart(SB), R10
	MOVV  $runtime·nodmaend(SB), R11
	SUB   R10, R11, R15  // size of non-DMA memory
	ADD   R13, R15, R14  // size of the whole free memory

	// we prefer the non-DMA memory for non-heap objects to preserve as much as
	// possible of the DMA capable memory for heap allocations
	SUB   R15, R12

	// reduce the arena by the remain of the non-heap space that did not fit in
	// the non-DMA memory, properly align the arena
	BLTZ  R12, 2(PC)
	SUB   R12, R13
	AND   $~(const_heapArenaBytes-1), R13
	SUB   R13, R9
	MOVV  R9, R12

	// save {free.start,free.end,nodma.start,nodma.end,arenaStart,arenaSize,size}
	MOVV  $runtime·noosMem(SB), R15
	MOVV  R8, 0(R15)
	MOVV  R9, 8(R15)
	MOVV  R10, 16(R15)
	MOVV  R11, 24(R15)
	MOVV  R12, 32(R15)
	MOVV  R13, 40(R15)
	MOVV  R14, 48(R15)

	// initialize noos tasker and Go scheduler
	JAL  runtime·taskerinit(SB)
	JAL  runtime·schedinit(SB)

	// allocate g0 for m0 and leave gh
	SUB   $16, R29
	MOVW  $(2*const_stackMin), R4
	MOVW  R4, 8(R29)
	JAL   runtime·malg(SB)
	MOVV  16(R29), R8  // newg in R8
	ADD   $16, R29

	// stackguard check during newproc requires valid stackguard1 but
	// malg sets it to 0xFFFFFFFF (mstart fixes this but is called later)
	MOVV  g_stackguard0(R8), R9
	MOVV  R9, g_stackguard1(R8)

	MOVV  $runtime·m0(SB), R9
	MOVV  R8, m_g0(R9)  // m0.g0 = newg
	MOVV  R9, g_m(R8)   // newg.m = m0

	// newg stack pointer to R29
	MOVV  (g_stack+stack_hi)(R8), R29

	// newg to g
	MOVV  g, R10
	MOVV  R8, g

	// fix cpu0.gh, cpu0.mh
	ADD   $cpuctx_mh, R10, R9  // R10 points to cpu0 (and to cpu0.gh at the same time)
	MOVV  R10, m_g0(R9)        // cpu0.mh.g0 = cpu0.gh
	MOVV  R10, m_gsignal(R9)   // cpu0.mh.gsignal = cpu0.gh (to easily check for handler mode)
	MOVV  R9, g_m(R10)         // cpu0.gh.m = cpu0.mh

	// TODO switch to the user mode?

	// create a new goroutine to start program
	SUB	$16, R29
	MOVV	$runtime·mainPC(SB), R8
	MOVV	R8, 8(R29)	// arg 1: fn
	MOVV	$0, R8
	MOVV	R8, 0(R29)	// dummy LR
	JAL	runtime·newproc(SB)
	ADD	$16, R29	// pop args and LR

	// enable interrupts
	// TODO where to enable interupts correctly?
	MOVW R0, M(C0_COMPARE)
	MOVW M(C0_SR), R8
	OR  $(SR_IE|INTR_SW|INTR_EXT), R8
	MOVW R8, M(C0_SR)

	// start this M
	JAL  runtime·mstart(SB)

	UNDEF
