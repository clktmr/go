// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !amd64,!386,!arm,!s390x,!arm64,!thumb

package sha1

var block = blockGeneric
