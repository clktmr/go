// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"runtime/internal/atomic"
	"unsafe"
)

// Tasker implements simple threads for noos target. It works directly on the m
// type pointers obtained from the Go scheduler.
//
// Every CPU has two local lists of threads: the cpuctx.runnable queue and the
// cpuctx.waitingt sorted list.
//
// The cpuctx.runnable queue contains runnable threads waiting for their
// timeslots on the CPU.
//
// The cpuctx.waitingt list contains threads that sleep until some time in the
// future, sorted by this time.
//
// Both lists are also accessed by other CPUs if they wakeup this CPU's thread
// that sleeps on a futex. The other CPU can only remove from the
// cpuctx.waitingt or add to the cpuctx.runnable.
//
// There is also tasker.waitingf, a global hash table of threads that sleep on
// futexes.
//
// The sleeping thread can be found in the cpuctx.waitingt, in the
// tasker.waitingf or in the both places.
//
// Tasker does not define new fields in m but reuses unused ones:
//
// - tls[0:4], ncgocall, thread: used by mq, msl, mcl types (see mq.go),
// - tls[4:6] and some other unused fields are used by architecture-specific code
//   (see: tasker_GOARCH.go),
// - cgoCallersUse, caughtsig : used by tasker.
//
// Tasker relies on the following architecture-specific functions:
//
// taskerpreinit, taskerinit
//
// These functions should initialize the tasker. Taskerpreinit is run before the
// Go scheduler and memory allocator are ready. It should setup a minimal
// enviroment for them.
//
// curcpuSleep, cpuctx.wakeup
//
// These functions are called to put the CPU to sleep and wake up it. If the
// CPU is running the wakeup command must be registered so the subsequent sleep
// will not happen (to avoid race condition). It is allowed that both of these
// functions do nothing.
//
// curcpuSavectxCall, curcpuSavectxSched
//
// This function is called to save the remaining context, not saved at syscall
// entry (eg. it can save FPU state).
//
// archnewm
//
// This function is called to create the inintial state of the new thread and
// save it in provided m.
//
// curcpuSchedule
//
// Run scheduler immediately or at syscall exit.
//
// The actual context switch is performed by architecture specific code at
// curcpuRunScheduler exit. It should check the cpuctx.newexe variable and if
// true switch the context to the new thread specified in cpuctx.exe.
//
// Tasker code does not use FPU so the architecture specific context switch
// code can avoid saving/restoring FPU context if not need.

func dummyNanotime() int64        { return 1 }
func dummySetalarm(ns int64) bool { return false }

var thetasker = tasker{nanotime: dummyNanotime, setalarm: dummySetalarm}

const fbnum = 4 // number of futex hash table buckets, must be power of two

type cpuctx struct {
	gh       g               // for ISRs, must be the first field in this struct
	t        *tasker         // points to thetasker
	exe      muintptr        // m currently executed by CPU
	newexe   bool            // for architecture-dependent code: exe changed
	runnable mq              // threads in runnable state
	waitingt msl             // threads waiting until some time elapses
	wakerq   [fbnum]notelist // futex wakeup request from interrupt handlers
	mh       m               // for ISRs, mostly not written so works as cache line pad
}

func (cpu *cpuctx) id() int { return int(cpu.gh.goid) }

type tasker struct {
	allcpu   []*cpuctx
	waitingf [fbnum]mcl // threads waiting on futex
	tidgen   uintptr

	timestart struct {
		sec  int64
		nsec int32
		mx   cpumtx
	}

	nanotime func() int64
	setalarm func(ns int64) bool

	newnanotime func() int64        // see embedded/rtos.SetSystemTimer
	newsetalarm func(ns int64) bool // see embedded/rtos.SetSystemTimer
}

//go:nosplit
func fhash(addr uintptr) int { return int(addr>>3) & (fbnum - 1) }

//go:nosplit
func (t *tasker) fbucketbyaddr(addr uintptr) *mcl {
	return &t.waitingf[fhash(addr)]
}

// gh is the first field of the cpuctx so we can benefit from the getg which is
// intrinsic function, often compiled to 0 or 1 instruction. Don't call in
// thread mode (valid only in handler mode).
//go:nosplit
func curcpu() *cpuctx { return (*cpuctx)(unsafe.Pointer(getg())) }

//go:nosplit
func taskerSetrunnable(m *m) {
	curcpu := curcpu()
	allcpu := curcpu.t.allcpu
	var (
		bestcpu *cpuctx
		bestn   int
	)
	p := m.nextp
	if p != 0 {
		goto byid
	}
	p = m.p
	if p != 0 {
		goto byid
	}
	p = m.oldp
	if p != 0 {
		goto byid
	}
	// naive search for the less loaded cpu
	bestcpu = curcpu
	bestn = bestcpu.runnable.atomicLen()
	for _, cpu := range allcpu {
		if n := cpu.runnable.atomicLen(); n < bestn {
			bestcpu = cpu
			bestn = n
		}
	}
	goto end
byid:
	bestcpu = allcpu[int(p.ptr().id)%len(allcpu)]
end:
	bestcpu.runnable.lock()
	n := bestcpu.runnable.n
	bestcpu.runnable.push(m)
	bestcpu.runnable.unlock()
	if bestcpu != curcpu && n == 0 {
		bestcpu.wakeup()
	}
}

//go:nosplit
func taskerFutexwakeup(fb *mcl, addr *uint32, cnt uint32) {
	for ; cnt != 0; cnt-- {
		fb.lock()
		m := fb.find(uintptr(unsafe.Pointer(addr)))
		if m == nil {
			fb.unlock()
			break
		}
		owned := true
		wt := mwt(m)
		if wt != nil {
			// this thread sleeps also in the cpuctx.waitingt
			owned = atomic.Cas(mownedptr(m), 0, 1)
		}
		if owned {
			fb.remove(m)
		}
		fb.unlock()
		if owned {
			if wt != nil {
				wt.lock()
				wt.remove(m)
				wt.unlock()
			}
			taskerSetrunnable(m)
		}
	}
	return
}

//go:nowritebarrierrec
//go:nosplit
func curcpuRunScheduler() {
	curcpu := curcpu()
	exe := curcpu.exe.ptr()
	for {
		// handle the wakeup requests from interrupt handlers
		for i := range curcpu.wakerq {
			n := curcpu.wakerq[i].removeall()
			for n != nil {
				next := n.release()
				taskerFutexwakeup(&curcpu.t.waitingf[i], key32(&n.key), 1)
				n = next
			}
		}

		var sleepuntil int64

		// waking up the threads sleeping in the curcpu.waitingt
		now := curcpu.t.nanotime()
		wt := &curcpu.waitingt
		for {
			wt.lock()
			m := wt.first()
			if m == nil {
				sleepuntil = -1
				break
			}
			sleepuntil = mval(m)
			if sleepuntil > now {
				break
			}
			owned := true
			addr := mkey(m)
			if addr != 0 {
				// this thread sleeps also in the tasker.waitingf
				owned = atomic.Cas(mownedptr(m), 0, 1)
			}
			if owned {
				wt.remove(m)
			}
			wt.unlock()
			if owned {
				if addr != 0 {
					fb := curcpu.t.fbucketbyaddr(addr)
					fb.lock()
					fb.remove(m)
					fb.unlock()
				}
				taskerSetrunnable(m)
			}
		}
		wt.unlock()

		// schedule the next thread from the curcpu.runnable
		curcpu.runnable.lock()
		next := curcpu.runnable.pop()
		if next != nil && exe != nil {
			curcpuSavectxSched()
			curcpu.runnable.push(exe)
		}
		n := curcpu.runnable.n
		curcpu.runnable.unlock()
		if next != nil {
			curcpu.exe.set(next)
			curcpu.newexe = true
			if n != 0 {
				//curcpu.t.setalarm(now + 2e6)
			}
			return
		}
		if exe != nil {
			return
		}

		// Nothing to execute. If this will be a work-stealing scheduler it will
		// try to steal some work from some other CPU here.
		if curcpu.t.setalarm(sleepuntil) {
			curcpuSleep()
		}
	}
}

//go:nowritebarrierrec
//go:nosplit
func rtos_notewakeup(n *notel) {
	if !atomic.Cas(key32(&n.key), 0, 1) {
		return
	}
	if !isr() {
		futexwakeup(key32(&n.key), 1)
		return
	}
	if n.acquire() {
		curcpu().wakerq[fhash(uintptr(unsafe.Pointer(&n.key)))].insert(n)
		curcpuWakeup()
	}
}

// notelist

// notel is a note that contains a link field to construct linked lists of notes
type notel struct {
	key  uintptr // must be the first field
	link notelptr
}

//go:nosplit
func (n *notel) acquire() bool {
	return (&n.link).atomicCAS(0, 1)
}

//go:nosplit
func (n *notel) release() *notel {
	next := n.link
	atomic.Storeuintptr((*uintptr)(&n.link), 0)
	return next.ptr()
}

//go:nosplit
func (n *notel) notelptr() notelptr { return notelptr(unsafe.Pointer(n)) }

type notelptr uintptr

//go:nosplit
func (n notelptr) ptr() *notel { return (*notel)(unsafe.Pointer(n)) }

//go:nosplit
func (p *notelptr) atomicLoad() notelptr {
	return notelptr(atomic.Loaduintptr((*uintptr)(p)))
}

//go:nosplit
func (p *notelptr) atomicCAS(old, new notelptr) bool {
	return atomic.Casuintptr((*uintptr)(p), uintptr(old), uintptr(new))
}

type notelist struct {
	head notelptr
}

// insert inserts n at the beginning of l. You must acquire n before insert it.
//go:nosplit
func (l *notelist) insert(n *notel) {
	if n.link != 1 {
		for {
			breakpoint()
		}
	}
	for {
		head := (&l.head).atomicLoad()
		n.link = head
		if (&l.head).atomicCAS(head, n.notelptr()) {
			return
		}
	}
}

// removeall removes and returns the whole content of l.
//go:nosplit
func (l *notelist) removeall() *notel {
	for {
		head := (&l.head).atomicLoad()
		if (&l.head).atomicCAS(head, 0) {
			return head.ptr()
		}
	}
}

// syscall handlers

//go:nowritebarrierrec
//go:nosplit
func syssetsystim1() {
	t := curcpu().t
	const n = unsafe.Sizeof(t.nanotime) / unsafe.Sizeof(uintptr(0))
	*(*[n]uintptr)(unsafe.Pointer(&t.nanotime)) = *(*[n]uintptr)(unsafe.Pointer(&t.newnanotime))
	*(*[n]uintptr)(unsafe.Pointer(&t.setalarm)) = *(*[n]uintptr)(unsafe.Pointer(&t.newsetalarm))
	curcpuSchedule() // ensure scheduler uses new timer
}

//go:nowritebarrierrec
//go:nosplit
func sysnanotime() int64 {
	return curcpu().t.nanotime()
}

//go:nowritebarrierrec
//go:nosplit
func sysnewosproc(m *m) {
	curcpu := curcpu()
	m.procid = uint64(atomic.Xadduintptr(&curcpu.t.tidgen, 1))
	archnewm(m)
	taskerSetrunnable(m)
}

//go:nowritebarrierrec
//go:nosplit
func sysexitThread(wait *uint32) {
	curcpu().exe = 0
	*wait = 0
	curcpuSchedule()
}

//go:nowritebarrierrec
//go:nosplit
func sysfutexsleep(addr *uint32, val uint32, ns int64) {
	if uint64(ns) < 64 {
		return // to short to sleep (64 ns selected arbitrary)
	}
	curcpu := curcpu()
	m := curcpu.exe.ptr()
	if ns >= 0 {
		// pre-insert m into curcpu.waitingt, m is not visible for other CPUs
		// until it is published in the thetasker.waitingf.
		msetval(m, curcpu.t.nanotime()+ns)
		msetowned(m, 0)
		wt := &curcpu.waitingt
		msetwt(m, wt)
		wt.lock()
		wt.insertbyval(m)
		wt.unlock()
	} else {
		msetwt(m, nil)
	}
	fb := curcpu.t.fbucketbyaddr(uintptr(unsafe.Pointer(addr)))
	fb.lock()
	sleep := (*addr == val)
	if sleep {
		curcpuSavectxCall()
		curcpu.exe = 0
		msetkey(m, uintptr(unsafe.Pointer(addr)))
		fb.push(m)
	}
	fb.unlock()
	if sleep {
		curcpuSchedule()
	} else if ns >= 0 {
		// revert the pre-insert
		curcpu.waitingt.lock()
		curcpu.waitingt.remove(m)
		curcpu.waitingt.unlock()
	}
}

//go:nowritebarrierrec
//go:nosplit
func sysfutexwakeup(addr *uint32, cnt uint32) {
	fb := curcpu().t.fbucketbyaddr(uintptr(unsafe.Pointer(addr)))
	taskerFutexwakeup(fb, addr, cnt)
}

//go:nowritebarrierrec
//go:nosplit
func sysnanosleep(ns int64) {
	if uint64(ns) < 64 {
		return // to short to sleep (64 ns selected arbitrary)
	}
	curcpuSavectxCall()
	curcpu := curcpu()
	m := curcpu.exe.ptr()
	curcpu.exe = 0
	msetkey(m, 0)
	msetval(m, curcpu.t.nanotime()+ns)
	wt := &curcpu.waitingt
	wt.lock()
	wt.insertbyval(m)
	wt.unlock()
	curcpuSchedule()
}

//go:nowritebarrierrec
//go:nosplit
func syswalltime() (sec int64, nsec int32) {
	t := curcpu().t
	t.timestart.mx.lock()
	sec = t.timestart.sec
	nsec = t.timestart.nsec
	t.timestart.mx.unlock()
	now := t.nanotime()
	s := now / 1e9
	ns := int32(now - s*1e9)
	sec += s
	nsec += ns
	if nsec >= 1e9 {
		sec++
		nsec -= 1e9
	}
	return
}

//go:nowritebarrierrec
//go:nosplit
func syssetwalltime(sec int64, nsec int32) {
	t := curcpu().t
	now := t.nanotime()
	s := now / 1e9
	ns := int32(now - s*1e9)
	sec -= s
	nsec -= ns
	if nsec < 0 {
		sec--
		nsec += 1e9
	}
	t.timestart.mx.lock()
	t.timestart.sec = sec
	t.timestart.nsec = nsec
	t.timestart.mx.unlock()
}

// m fields used

//go:nosplit
func mownedptr(m *m) *uint32 { return &m.cgoCallersUse }

//go:nosplit
func msetowned(m *m, owned uint32) { m.cgoCallersUse = owned }

//go:nosplit
func mwt(m *m) *msl { return (*msl)(unsafe.Pointer(m.caughtsig)) }

//go:nosplit
func msetwt(m *m, wt *msl) { m.caughtsig = guintptr(unsafe.Pointer(wt)) }
