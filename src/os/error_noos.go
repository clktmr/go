// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import "syscall"

type syscallErrorType = *syscall.Error

var errENOSYS = syscall.NewError("function not implemented")
var errERANGE = syscall.NewError("out of range")
var errENOMEM = syscall.NewError("cannot allocate memory")
