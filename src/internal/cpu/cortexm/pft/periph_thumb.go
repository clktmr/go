// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Instances:
//  PFT  0xE000ED78  -  -  Processor features registers
// Registers:
//  0x00 32  CLIDR   Cache Level ID
//  0x00 32  CTR     Cache Type
//  0x00 32  CCSIDR  Cache Size ID
//  0x00 32  CSSELR  Cache Size Selection
package pft

const (
	CL1I  CLIDR = 1 << 0    //+ Instruction L1 cache implemented.
	CL1D  CLIDR = 1 << 1    //+ Data cache L1 implemented.
	CL1U  CLIDR = 1 << 2    //+ Unified L1 cache.
	CL2I  CLIDR = 1 << 3    //+ Instruction L2 cache implemented.
	CL2D  CLIDR = 1 << 4    //+ Data cache L2 implemented.
	CL2U  CLIDR = 1 << 5    //+ Unified L2 cache.
	CL3I  CLIDR = 1 << 6    //+ Instruction L3 cache implemented.
	CL3D  CLIDR = 1 << 7    //+ Data cache L3 implemented.
	CL3U  CLIDR = 1 << 8    //+ Unified L3 cache.
	CL4I  CLIDR = 1 << 9    //+ Instruction L4 cache implemented.
	CL4D  CLIDR = 1 << 10   //+ Data cache L4 implemented.
	CL4U  CLIDR = 1 << 11   //+ Unified L4 cache.
	CL5I  CLIDR = 1 << 12   //+ Instruction L5 cache implemented.
	CL5D  CLIDR = 1 << 13   //+ Data cache L5 implemented.
	CL5U  CLIDR = 1 << 14   //+ Unified L5 cache.
	CL6I  CLIDR = 1 << 15   //+ Instruction L6 cache implemented.
	CL6D  CLIDR = 1 << 16   //+ Data cache L6 implemented.
	CL6U  CLIDR = 1 << 17   //+ Unified L6 cache.
	CL7I  CLIDR = 1 << 18   //+ Instruction L7 cache implemented.
	CL7D  CLIDR = 1 << 19   //+ Data cache L7 implemented.
	CL7U  CLIDR = 1 << 20   //+ Unified L7 cache.
	LoUIS CLIDR = 0x7 << 21 //+
	LoC   CLIDR = 0x7 << 24 //+ Level of Coherency.
	LoU   CLIDR = 0x7 << 27 //+ Level of Unification.
)

const (
	CL1In  = 0
	CL1Dn  = 1
	CL1Un  = 2
	CL2In  = 3
	CL2Dn  = 4
	CL2Un  = 5
	CL3In  = 6
	CL3Dn  = 7
	CL3Un  = 8
	CL4In  = 9
	CL4Dn  = 10
	CL4Un  = 11
	CL5In  = 12
	CL5Dn  = 13
	CL5Un  = 14
	CL6In  = 15
	CL6Dn  = 16
	CL6Un  = 17
	CL7In  = 18
	CL7Dn  = 19
	CL7Un  = 20
	LoUISn = 21
	LoCn   = 24
	LoUn   = 27
)

const (
	IMinLine CTR = 0xf << 0  //+ Smallest cache line of all the I-caches.
	DMinLine CTR = 0xf << 16 //+ Smallest cache line of all the D/U-caches.
	ERG      CTR = 0xf << 20 //+ Exclusives Reservation Granule.
	CWG      CTR = 0xf << 24 //+ Cache Writeback Granule.
	Format   CTR = 0x7 << 29 //+ Register format (4: ARMv7 format).
)

const (
	IMinLinen = 0
	DMinLinen = 16
	ERGn      = 20
	CWGn      = 24
	Formatn   = 29
)

const (
	LineSize      CCSIDR = 0x7 << 0     //+ Number of words in cache line (log2(n)-2).
	Associativity CCSIDR = 0x3ff << 3   //+ Number of ways - 1.
	NumSets       CCSIDR = 0x7fff << 13 //+ Number of sets - 1.
	WA            CCSIDR = 1 << 28      //+ Write allocation support.
	RA            CCSIDR = 1 << 29      //+ Read allocation support.
	WB            CCSIDR = 1 << 30      //+ Write-Back support.
	WT            CCSIDR = 1 << 31      //+ Write-Through support.
)

const (
	LineSizen      = 0
	Associativityn = 3
	NumSetsn       = 13
	WAn            = 28
	RAn            = 29
	WBn            = 30
	WTn            = 31
)

const (
	InD   CSSELR = 1 << 0   //+ Selection of 1:instruction or 0:data cache.
	Level CSSELR = 0x7 << 1 //+ Cache level selected (0: level1).
)

const (
	InDn   = 0
	Leveln = 1
)
