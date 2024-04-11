// go run linux/mksysnum.go -Wall -Werror -static -I/tmp/include /tmp/include/asm/unistd.h
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build thumb,linux

package unix

const (
	SYS_RESTART_SYSCALL        = 0
	SYS_EXIT                   = 1
	SYS_FORK                   = 2
	SYS_READ                   = 3
	SYS_WRITE                  = 4
	SYS_OPEN                   = 5
	SYS_CLOSE                  = 6
	SYS_CREAT                  = 8
	SYS_LINK                   = 9
	SYS_UNLINK                 = 10
	SYS_EXECVE                 = 11
	SYS_CHDIR                  = 12
	SYS_MKNOD                  = 14
	SYS_CHMOD                  = 15
	SYS_LCHOWN                 = 16
	SYS_LSEEK                  = 19
	SYS_GETPID                 = 20
	SYS_MOUNT                  = 21
	SYS_SETUID                 = 23
	SYS_GETUID                 = 24
	SYS_PTRACE                 = 26
	SYS_PAUSE                  = 29
	SYS_ACCESS                 = 33
	SYS_NICE                   = 34
	SYS_SYNC                   = 36
	SYS_KILL                   = 37
	SYS_RENAME                 = 38
	SYS_MKDIR                  = 39
	SYS_RMDIR                  = 40
	SYS_DUP                    = 41
	SYS_PIPE                   = 42
	SYS_TIMES                  = 43
	SYS_BRK                    = 45
	SYS_SETGID                 = 46
	SYS_GETGID                 = 47
	SYS_GETEUID                = 49
	SYS_GETEGID                = 50
	SYS_ACCT                   = 51
	SYS_UMOUNT2                = 52
	SYS_IOCTL                  = 54
	SYS_FCNTL                  = 55
	SYS_SETPGID                = 57
	SYS_UMASK                  = 60
	SYS_CHROOT                 = 61
	SYS_USTAT                  = 62
	SYS_DUP2                   = 63
	SYS_GETPPID                = 64
	SYS_GETPGRP                = 65
	SYS_SETSID                 = 66
	SYS_SIGACTION              = 67
	SYS_SETREUID               = 70
	SYS_SETREGID               = 71
	SYS_SIGSUSPEND             = 72
	SYS_SIGPENDING             = 73
	SYS_SETHOSTNAME            = 74
	SYS_SETRLIMIT              = 75
	SYS_GETRUSAGE              = 77
	SYS_GETTIMEOFDAY           = 78
	SYS_SETTIMEOFDAY           = 79
	SYS_GETGROUPS              = 80
	SYS_SETGROUPS              = 81
	SYS_SYMLINK                = 83
	SYS_READLINK               = 85
	SYS_USELIB                 = 86
	SYS_SWAPON                 = 87
	SYS_REBOOT                 = 88
	SYS_MUNMAP                 = 91
	SYS_TRUNCATE               = 92
	SYS_FTRUNCATE              = 93
	SYS_FCHMOD                 = 94
	SYS_FCHOWN                 = 95
	SYS_GETPRIORITY            = 96
	SYS_SETPRIORITY            = 97
	SYS_STATFS                 = 99
	SYS_FSTATFS                = 100
	SYS_SYSLOG                 = 103
	SYS_SETITIMER              = 104
	SYS_GETITIMER              = 105
	SYS_STAT                   = 106
	SYS_LSTAT                  = 107
	SYS_FSTAT                  = 108
	SYS_VHANGUP                = 111
	SYS_WAIT4                  = 114
	SYS_SWAPOFF                = 115
	SYS_SYSINFO                = 116
	SYS_FSYNC                  = 118
	SYS_SIGRETURN              = 119
	SYS_CLONE                  = 120
	SYS_SETDOMAINNAME          = 121
	SYS_UNAME                  = 122
	SYS_ADJTIMEX               = 124
	SYS_MPROTECT               = 125
	SYS_SIGPROCMASK            = 126
	SYS_INIT_MODULE            = 128
	SYS_DELETE_MODULE          = 129
	SYS_QUOTACTL               = 131
	SYS_GETPGID                = 132
	SYS_FCHDIR                 = 133
	SYS_BDFLUSH                = 134
	SYS_SYSFS                  = 135
	SYS_PERSONALITY            = 136
	SYS_SETFSUID               = 138
	SYS_SETFSGID               = 139
	SYS__LLSEEK                = 140
	SYS_GETDENTS               = 141
	SYS__NEWSELECT             = 142
	SYS_FLOCK                  = 143
	SYS_MSYNC                  = 144
	SYS_READV                  = 145
	SYS_WRITEV                 = 146
	SYS_GETSID                 = 147
	SYS_FDATASYNC              = 148
	SYS__SYSCTL                = 149
	SYS_MLOCK                  = 150
	SYS_MUNLOCK                = 151
	SYS_MLOCKALL               = 152
	SYS_MUNLOCKALL             = 153
	SYS_SCHED_SETPARAM         = 154
	SYS_SCHED_GETPARAM         = 155
	SYS_SCHED_SETSCHEDULER     = 156
	SYS_SCHED_GETSCHEDULER     = 157
	SYS_SCHED_YIELD            = 158
	SYS_SCHED_GET_PRIORITY_MAX = 159
	SYS_SCHED_GET_PRIORITY_MIN = 160
	SYS_SCHED_RR_GET_INTERVAL  = 161
	SYS_NANOSLEEP              = 162
	SYS_MREMAP                 = 163
	SYS_SETRESUID              = 164
	SYS_GETRESUID              = 165
	SYS_POLL                   = 168
	SYS_NFSSERVCTL             = 169
	SYS_SETRESGID              = 170
	SYS_GETRESGID              = 171
	SYS_PRCTL                  = 172
	SYS_RT_SIGRETURN           = 173
	SYS_RT_SIGACTION           = 174
	SYS_RT_SIGPROCMASK         = 175
	SYS_RT_SIGPENDING          = 176
	SYS_RT_SIGTIMEDWAIT        = 177
	SYS_RT_SIGQUEUEINFO        = 178
	SYS_RT_SIGSUSPEND          = 179
	SYS_PREAD64                = 180
	SYS_PWRITE64               = 181
	SYS_CHOWN                  = 182
	SYS_GETCWD                 = 183
	SYS_CAPGET                 = 184
	SYS_CAPSET                 = 185
	SYS_SIGALTSTACK            = 186
	SYS_SENDFILE               = 187
	SYS_VFORK                  = 190
	SYS_UGETRLIMIT             = 191
	SYS_MMAP2                  = 192
	SYS_TRUNCATE64             = 193
	SYS_FTRUNCATE64            = 194
	SYS_STAT64                 = 195
	SYS_LSTAT64                = 196
	SYS_FSTAT64                = 197
	SYS_LCHOWN32               = 198
	SYS_GETUID32               = 199
	SYS_GETGID32               = 200
	SYS_GETEUID32              = 201
	SYS_GETEGID32              = 202
	SYS_SETREUID32             = 203
	SYS_SETREGID32             = 204
	SYS_GETGROUPS32            = 205
	SYS_SETGROUPS32            = 206
	SYS_FCHOWN32               = 207
	SYS_SETRESUID32            = 208
	SYS_GETRESUID32            = 209
	SYS_SETRESGID32            = 210
	SYS_GETRESGID32            = 211
	SYS_CHOWN32                = 212
	SYS_SETUID32               = 213
	SYS_SETGID32               = 214
	SYS_SETFSUID32             = 215
	SYS_SETFSGID32             = 216
	SYS_GETDENTS64             = 217
	SYS_PIVOT_ROOT             = 218
	SYS_MINCORE                = 219
	SYS_MADVISE                = 220
	SYS_FCNTL64                = 221
	SYS_GETTID                 = 224
	SYS_READAHEAD              = 225
	SYS_SETXATTR               = 226
	SYS_LSETXATTR              = 227
	SYS_FSETXATTR              = 228
	SYS_GETXATTR               = 229
	SYS_LGETXATTR              = 230
	SYS_FGETXATTR              = 231
	SYS_LISTXATTR              = 232
	SYS_LLISTXATTR             = 233
	SYS_FLISTXATTR             = 234
	SYS_REMOVEXATTR            = 235
	SYS_LREMOVEXATTR           = 236
	SYS_FREMOVEXATTR           = 237
	SYS_TKILL                  = 238
	SYS_SENDFILE64             = 239
	SYS_FUTEX                  = 240
	SYS_SCHED_SETAFFINITY      = 241
	SYS_SCHED_GETAFFINITY      = 242
	SYS_IO_SETUP               = 243
	SYS_IO_DESTROY             = 244
	SYS_IO_GETEVENTS           = 245
	SYS_IO_SUBMIT              = 246
	SYS_IO_CANCEL              = 247
	SYS_EXIT_GROUP             = 248
	SYS_LOOKUP_DCOOKIE         = 249
	SYS_EPOLL_CREATE           = 250
	SYS_EPOLL_CTL              = 251
	SYS_EPOLL_WAIT             = 252
	SYS_REMAP_FILE_PAGES       = 253
	SYS_SET_TID_ADDRESS        = 256
	SYS_TIMER_CREATE           = 257
	SYS_TIMER_SETTIME          = 258
	SYS_TIMER_GETTIME          = 259
	SYS_TIMER_GETOVERRUN       = 260
	SYS_TIMER_DELETE           = 261
	SYS_CLOCK_SETTIME          = 262
	SYS_CLOCK_GETTIME          = 263
	SYS_CLOCK_GETRES           = 264
	SYS_CLOCK_NANOSLEEP        = 265
	SYS_STATFS64               = 266
	SYS_FSTATFS64              = 267
	SYS_TGKILL                 = 268
	SYS_UTIMES                 = 269
	SYS_ARM_FADVISE64_64       = 270
	SYS_PCICONFIG_IOBASE       = 271
	SYS_PCICONFIG_READ         = 272
	SYS_PCICONFIG_WRITE        = 273
	SYS_MQ_OPEN                = 274
	SYS_MQ_UNLINK              = 275
	SYS_MQ_TIMEDSEND           = 276
	SYS_MQ_TIMEDRECEIVE        = 277
	SYS_MQ_NOTIFY              = 278
	SYS_MQ_GETSETATTR          = 279
	SYS_WAITID                 = 280
	SYS_SOCKET                 = 281
	SYS_BIND                   = 282
	SYS_CONNECT                = 283
	SYS_LISTEN                 = 284
	SYS_ACCEPT                 = 285
	SYS_GETSOCKNAME            = 286
	SYS_GETPEERNAME            = 287
	SYS_SOCKETPAIR             = 288
	SYS_SEND                   = 289
	SYS_SENDTO                 = 290
	SYS_RECV                   = 291
	SYS_RECVFROM               = 292
	SYS_SHUTDOWN               = 293
	SYS_SETSOCKOPT             = 294
	SYS_GETSOCKOPT             = 295
	SYS_SENDMSG                = 296
	SYS_RECVMSG                = 297
	SYS_SEMOP                  = 298
	SYS_SEMGET                 = 299
	SYS_SEMCTL                 = 300
	SYS_MSGSND                 = 301
	SYS_MSGRCV                 = 302
	SYS_MSGGET                 = 303
	SYS_MSGCTL                 = 304
	SYS_SHMAT                  = 305
	SYS_SHMDT                  = 306
	SYS_SHMGET                 = 307
	SYS_SHMCTL                 = 308
	SYS_ADD_KEY                = 309
	SYS_REQUEST_KEY            = 310
	SYS_KEYCTL                 = 311
	SYS_SEMTIMEDOP             = 312
	SYS_VSERVER                = 313
	SYS_IOPRIO_SET             = 314
	SYS_IOPRIO_GET             = 315
	SYS_INOTIFY_INIT           = 316
	SYS_INOTIFY_ADD_WATCH      = 317
	SYS_INOTIFY_RM_WATCH       = 318
	SYS_MBIND                  = 319
	SYS_GET_MEMPOLICY          = 320
	SYS_SET_MEMPOLICY          = 321
	SYS_OPENAT                 = 322
	SYS_MKDIRAT                = 323
	SYS_MKNODAT                = 324
	SYS_FCHOWNAT               = 325
	SYS_FUTIMESAT              = 326
	SYS_FSTATAT64              = 327
	SYS_UNLINKAT               = 328
	SYS_RENAMEAT               = 329
	SYS_LINKAT                 = 330
	SYS_SYMLINKAT              = 331
	SYS_READLINKAT             = 332
	SYS_FCHMODAT               = 333
	SYS_FACCESSAT              = 334
	SYS_PSELECT6               = 335
	SYS_PPOLL                  = 336
	SYS_UNSHARE                = 337
	SYS_SET_ROBUST_LIST        = 338
	SYS_GET_ROBUST_LIST        = 339
	SYS_SPLICE                 = 340
	SYS_ARM_SYNC_FILE_RANGE    = 341
	SYS_TEE                    = 342
	SYS_VMSPLICE               = 343
	SYS_MOVE_PAGES             = 344
	SYS_GETCPU                 = 345
	SYS_EPOLL_PWAIT            = 346
	SYS_KEXEC_LOAD             = 347
	SYS_UTIMENSAT              = 348
	SYS_SIGNALFD               = 349
	SYS_TIMERFD_CREATE         = 350
	SYS_EVENTFD                = 351
	SYS_FALLOCATE              = 352
	SYS_TIMERFD_SETTIME        = 353
	SYS_TIMERFD_GETTIME        = 354
	SYS_SIGNALFD4              = 355
	SYS_EVENTFD2               = 356
	SYS_EPOLL_CREATE1          = 357
	SYS_DUP3                   = 358
	SYS_PIPE2                  = 359
	SYS_INOTIFY_INIT1          = 360
	SYS_PREADV                 = 361
	SYS_PWRITEV                = 362
	SYS_RT_TGSIGQUEUEINFO      = 363
	SYS_PERF_EVENT_OPEN        = 364
	SYS_RECVMMSG               = 365
	SYS_ACCEPT4                = 366
	SYS_FANOTIFY_INIT          = 367
	SYS_FANOTIFY_MARK          = 368
	SYS_PRLIMIT64              = 369
	SYS_NAME_TO_HANDLE_AT      = 370
	SYS_OPEN_BY_HANDLE_AT      = 371
	SYS_CLOCK_ADJTIME          = 372
	SYS_SYNCFS                 = 373
	SYS_SENDMMSG               = 374
	SYS_SETNS                  = 375
	SYS_PROCESS_VM_READV       = 376
	SYS_PROCESS_VM_WRITEV      = 377
	SYS_KCMP                   = 378
	SYS_FINIT_MODULE           = 379
	SYS_SCHED_SETATTR          = 380
	SYS_SCHED_GETATTR          = 381
	SYS_RENAMEAT2              = 382
	SYS_SECCOMP                = 383
	SYS_GETRANDOM              = 384
	SYS_MEMFD_CREATE           = 385
	SYS_BPF                    = 386
	SYS_EXECVEAT               = 387
	SYS_USERFAULTFD            = 388
	SYS_MEMBARRIER             = 389
	SYS_MLOCK2                 = 390
	SYS_COPY_FILE_RANGE        = 391
	SYS_PREADV2                = 392
	SYS_PWRITEV2               = 393
	SYS_PKEY_MPROTECT          = 394
	SYS_PKEY_ALLOC             = 395
	SYS_PKEY_FREE              = 396
	SYS_STATX                  = 397
	SYS_RSEQ                   = 398
	SYS_IO_PGETEVENTS          = 399
)
