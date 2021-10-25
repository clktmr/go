// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "go_tls.h"
#include "textflag.h"
#include "syscall_noos.h"

// This code uses ADD and ORR instructions when wants to set bits from 0 to 1.  // Mixing these two ways to do the same thing may seem seemingly inconsistent
// but it is not. The shorter encoding is prefered. If both gives the same
// length the ORR instruction is used because of its less energy per instruction
// factor (see: https://www.ics.forth.gr/carv/greenvm/files/tr450.pdf).

// TODO: Use ICSR.RETTOBASE to avoid manual stacking of previously stacked
// registers.

#define ICSR_ADDR 0xE000ED04
#define ICSR_PENDSVCLR (1<<27)

// identcurcpu indetifies the current CPU and returns a pointer to its cpuctx in
// R0. It can clobber R0-R4,LR registers (other registers must be preserved).
TEXT runtime·identcurcpu(SB),NOSPLIT|NOFRAME,$0-0
	// for now only single CPU is supported (see also cpuid, osinit)
	MOVW  $·cpu0(SB), R0
	RET

// func sev()
TEXT ·sev(SB),NOSPLIT|NOFRAME,$0-0
	SEV
	RET

// func cpucontrol() uint32
TEXT ·cpucontrol(SB),NOSPLIT|NOFRAME,$0-4
	MOVW  CONTROL, R0
	MOVW  R0, ret+0(FP)
	RET

// func setcpucontrol(ctrl uint32)
TEXT ·setcpucontrol(SB),NOSPLIT|NOFRAME,$0-4
	MOVW  ctrl+0(FP), R0
	MOVW  R0, CONTROL
	RET

// func curcpuSleep()
TEXT ·curcpuSleep(SB),NOSPLIT|NOFRAME,$0-0
	DSB  // flush CPU write buffers before sleep
	WFE
	// still in pendsvHandler so clear PendSV to avoid unnecessary reentry
	MOVW  $ICSR_ADDR, R0
	MOVW  $ICSR_PENDSVCLR, R1
	MOVW  R1, (R0)
	DMB   // ensure clearing happens before reading something for what we woke
	RET

// Exception handlers

TEXT runtime·faultHandler(SB),NOSPLIT|NOFRAME,$0-0
	// At this point a lot of things can be broken so don't touch
	// stack nor memory. Do only few things that helps debuging.
	TST   $4, LR
	BNE   3(PC)
	MOVW  MSP, R1
	B     2(PC)
	MOVW  PSP, R1
	MOVW  IPSR, R0

	// Now R0 and R1 contain useful information.

	// R0 contains exception number:
	// 3: HardFault  - see HFSR: x/xw 0xE000ED2C
	// 4: MemManage  - see MMSR: x/xb 0xE000ED28, MMAR: x/xw 0xE000ED34
	// 5: BusFault   - see BFSR: x/xb 0xE000ED29, BFAR: x/xw 0xE000ED38
	// 6: UsageFault - see UFSR: x/xh 0xE000ED2A

	// R1 should contain pointer to the exception stack frame:
	// (R1) -> [R0, R1, R2, R3, IP, LR, PC, PSR]
	// If R1 points to the valid memory examine:
	// 1. Where PC points.
	// 2. Thumb bit in PSR
	// 3. IPSR in PSR

	// To print stack frame in GDB use x/8xw $r1

	// MemManage with MMSR = 0x82 and MMAR = 0 means probably the nil pointer
	// dereference (TODO: print stack trace). Useful GDB commands:
	//
	//	l **($r1+24)
	//	b **($r1+24)

	BKPT
	B   -1(PC)

TEXT runtime·reservedHandler(SB),NOSPLIT|NOFRAME,$0-0
	B   ·faultHandler(SB)

TEXT runtime·nmiHandler(SB),NOSPLIT|NOFRAME,$0-0
	B   ·faultHandler(SB)

TEXT runtime·hardfaultHandler(SB),NOSPLIT|NOFRAME,$0-0
	B   ·faultHandler(SB)

TEXT runtime·memmanageHandler(SB),NOSPLIT|NOFRAME,$0-0
	B   ·faultHandler(SB)

TEXT runtime·busfaultHandler(SB),NOSPLIT|NOFRAME,$0-0
	B   ·faultHandler(SB)

TEXT runtime·usagefaultHandler(SB),NOSPLIT|NOFRAME,$0-0
	B   ·faultHandler(SB)

TEXT runtime·securefaultHandler(SB),NOSPLIT|NOFRAME,$0-0
	B   ·faultHandler(SB)

TEXT runtime·debugmonHandler(SB),NOSPLIT|NOFRAME,$0-0
	B   ·faultHandler(SB)


#define sysMaxArgs (24+4) // max. size of argumants and return values + LR


// svcallHandler handles synhronous SVCall exception generated by SWI (SVC)
// instruction. The tiny wrappers over SWI instruction add three additional
// parameters in registers:
//
// R4 - syscall number,
// R5 - argument data size on the stack (+4 for frame-pointer),
// R6 - return data size on the stack.
//
// The R0-R3,R12,LR,PC,PSR registers have been pushed on the stack at exception
// entry. If FP registers have been used the place for D0-D7,FPSCR have been
// also reserved on the stack (lazy stacking). The content of R0-R3,R12
// registers can be broken by late-arriving higher-priority exception so using
// R4-R6 avoids reading from the stack.
//
// The ABI0 does not follow the AAPCS: almost all GP registers and all FP
// registers are caller saved. Therefore other (asynchronous) excepion handlers
// need to save and restore not stacked GP registers and FP ones (if extended
// frame was used). The SVCall is synchronous exception so this handler is
// called almost like normal function and does not have to save any registers
// except g (R10) and LR (EXC_RETURN).
TEXT runtime·svcallHandler(SB),NOSPLIT|NOFRAME,$0-0
	// check the syscall number
	CMP  $SYS_NUM, R4
	BGE  badSyscall

	// stacked SP to R7
	TST      $(1<<2), LR
	MOVW.EQ  MSP, R7
	MOVW.NE  PSP, R7
	MOVW     R7, R1

	// check does paddnig was added to the frame (stacked xPSR bit 9)
	MOVW    (7*4)(R7), R0
	TST     $(1<<9), R0
	ADD.NE  $4, R1

	// check does the extended frame is used
	TST     $(1<<4), LR
	ADD.NE  $(8*4), R1
	ADD.EQ  $(26*4), R1

	// push g, LR on the stack
	MOVM.DB.W  [g,LR], (R13)

	// make space on the stack for arguments + 3 registers
	SUB  $(sysMaxArgs+3*4), R13

	// copy arguments from the caller's stack
	MOVW  $·duffcopy+1024(SB), R0
	MOVW  R13, R2
	SUB   R5<<1, R0
	MOVW  LR, R5  // save EXC_RETURN before call
	BL    (R0)

	// save data needed to copy the return values back to the caller's stack
	ADD        $sysMaxArgs, R13, R0
	MOVM.IA.W  [R1,R2,R6], (R0)

	// current CPU context to R0
	BL  ·identcurcpu(SB)

	// check for fast syscall (unfortunately it lets through the calls by
	// higher priority exceptions that are disallowed to use syscalls at all)
	CMP  $SYS_LAST_FAST, R4
	BLS  fast

	// check for syscall from interrupt handler
	CMP  R0, g
	BEQ  badHandlerCall  // syscall not allowed in handler mode

	// save thread context (small): SP+CONTROL[nPRIV], EXC_RETURN, g
	MOVW  (cpuctx_exe)(R0), R3
	MOVW  CONTROL, R2
	AND   $const_thrPrivLevel, R2
	ORR   R2, R7
	ADD   $const_thrSmallCtx, R7  // set thrSmallCtx (only g saved)
	MOVW  R7, (m_tls+const_msp*4)(R3)
	MOVW  R5, (m_tls+const_mer*4)(R3)
	MOVW  g, (m_libcall)(R3)

fast:
	// call the service routine
	MOVW  R0, g
	MOVW  $·syscalls(SB), R0
	MOVW  (R0)(R4*4), R0
	BL    (R0)

	// copy the return values back to the caller's stack
	MOVW  (sysMaxArgs+2*4)(R13), R5
	CBZ   R5, 6(PC)  // check if is something to copy
	MOVW  (sysMaxArgs+0*4)(R13), R2
	MOVW  (sysMaxArgs+1*4)(R13), R1
	MOVW  $·duffcopy+1024(SB), R0
	SUB   R5<<1, R0
	BL    (R0)

	// wind up the stack and return from syscall
	ADD        $(sysMaxArgs+3*4), R13
	MOVM.IA.W  (R13), [g,R15]

badSyscall:
	BKPT
	B   -1(PC)

badHandlerCall:
	BKPT
	B   -1(PC)


// pendsvHandler handles asynhronous PendSV exceptions generated by the system
// timer routine or system calls, to schedule/wakeup next thread.
TEXT runtime·pendsvHandler(SB),NOSPLIT|NOFRAME,$0-0
	// load cpuctx
	MOVW  LR, R12
	BL    ·identcurcpu(SB)  // current CPU context to R0

	// if cpuctx.schedule then context saved by syscall
	MOVBU  (cpuctx_schedule)(R0), R3
	CBNZ   R3, contextSaved

	// save not stacked registers (R4-R11), SP, CONTROL[nPRIV], EXC_RETURN
	MOVW     (cpuctx_exe)(R0), R3
	TST      $(1<<2), R12
	MOVW.EQ  MSP, R1
	MOVW.NE  PSP, R1
	MOVW     CONTROL, R2
	AND      $const_thrPrivLevel, R2
	ORR      R2, R1
	MOVW     R1, (m_tls+const_msp*4)(R3)
	MOVW     R12, (m_tls+const_mer*4)(R3)
	ADD      $m_libcall, R3
	MOVM.IA  [R4-R11], (R3)

contextSaved:
	MOVW  $0, R3
	MOVB  R3, (cpuctx_schedule)(R0)

	// clear PendSV if set again to avoid unnecessary reentry to this handler
	MOVW  $ICSR_ADDR, R1
	MOVW  $ICSR_PENDSVCLR, R2
	MOVW  R2, (R1)
	DMB   // ensure clearing happens before checking nanotime and futexes

	// enter scheduler
	MOVW  R0, g
	BL    ·curcpuRunScheduler(SB)

	// load SP+CONTROL[nPRIV], EXC_RETURN from new/old context pointed by exe
	MOVW  (cpuctx_exe)(g), R3
	MOVW  (m_tls+const_msp*4)(R3), R0
	MOVW  (m_tls+const_mer*4)(R3), R1

	// check does the context changed
	MOVBU  (cpuctx_newexe)(g), R2
	CBNZ   R2, newexe

	// fast path if exe did not changed (cpuctx.newexe == false)
	TST      $const_thrSmallCtx, R0
	MOVW.NE  (m_libcall)(R3), g
	B.NE     (R1)
	ADD      $m_libcall, R3
	MOVM.IA  (R3), [R4-R11]
	B        (R1)

newexe:
	// clear cpuctx.newexe
	MOVW  $0, R2
	MOVB  R2, (cpuctx_newexe)(g)

	// restore privilege level
	MOVW  CONTROL, R2
	BIC   $const_thrPrivLevel, R2
	AND   $const_thrPrivLevel, R0, R4
	ORR   R4, R2
	MOVW  R2, CONTROL

	// restore PSP or MSP
	BIC      $(const_thrPrivLevel+const_thrSmallCtx), R0, R2
	TST      $(1<<2), R1
	MOVW.EQ  R2, MSP
	MOVW.NE  R2, PSP

	// fast path in case of small context (only g saved in libcall)
	TST      $const_thrSmallCtx, R0
	MOVW.NE  (m_libcall)(R3), g
	B.NE     (R1)

	// restore registers saved in m.libcall
	ADD        $m_libcall, R3
	MOVM.IA.W  (R3), [R4-R11]
	TST        $0x10, R1
	BNE        3(PC)
	HWORD      $0xEC93  // VLDM R3
	HWORD      $0x8B10  // [D8-D15]
	B          (R1)

TEXT runtime·curcpuSavectxSched(SB),NOSPLIT|NOFRAME,$0-0
	MOVW  (cpuctx_exe)(g), R0

	MOVW  (m_tls+const_mer*4)(R0), R1
	TST   $0x10, R1
	RET.NE

	ADD   $(m_libcall+8*4), R0
	MOVW  CONTROL, R1
	CPSID
	HWORD  $0xEC80      // VSTM R0
	HWORD  $0x8B10      // [D8-D15]
	MOVW   R1, CONTROL  // to avoid stacking again by higher priority handler
	CPSIE
	RET

TEXT runtime·unhandledException(SB),NOSPLIT|NOFRAME,$0-0
	BKPT
	B   -1(PC)
