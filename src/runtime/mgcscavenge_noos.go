// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "runtime/internal/atomic"

var scavenge struct {
	gcPercentGoal  atomic.Uint64
	assistTime     atomic.Int64
	backgroundTime atomic.Int64
}

var scavenger scavengerState

type scavengerState struct {
	sysmonWake atomic.Uint32
}

func (s *scavengerState) ready() {}
func (s *scavengerState) wake()  {}

type scavengeIndex struct{}

func (s *scavengeIndex) mark(base, limit uintptr)                              {}
func (s *scavengeIndex) nextGen()                                              {}
func (s *scavengeIndex) init(test bool, sysStat *sysMemStat) uintptr           { return 0 }
func (s *scavengeIndex) grow(base, limit uintptr, sysStat *sysMemStat) uintptr { return 0 }
func (*scavengeIndex) alloc(ci chunkIdx, npages uint)                          {}
func (s *scavengeIndex) free(ci chunkIdx, page, npages uint)                   {}

func bgscavenge(c chan int)                                            {}
func gcPaceScavenger(memoryLimit int64, heapGoal, lastHeapGoal uint64) {}
func heapRetained() uint64                                             { return 0 }
func printScavTrace(releasedBg, releasedEager uintptr, forced bool)    {}

func (p *pageAlloc) scavenge(nbytes uintptr, shouldStop func() bool, force bool) uintptr { return 0 }
