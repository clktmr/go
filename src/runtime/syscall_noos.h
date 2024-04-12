// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Some syscalls are alowed to be called by low priority interrupt handlers.
//
// Other syscalls, that can't run concurently with the scheduler, are available
// only to threads.

#define SYS_nanotime       0
#define SYS_irqctl         1
#define SYS_setprivlevel   2
#define SYS_write          3

#define SYS_LAST_FAST      3

#define SYS_setsystim1     4
#define SYS_setsyswriter1  5
#define SYS_newosproc      6
#define SYS_exitThread     7
#define SYS_futexsleep     8
#define SYS_futexwakeup    9
#define SYS_osyield       10
#define SYS_nanosleep     11

#define SYS_NUM           12
