// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"context"
	"syscall"
)

func lookupProtocol(ctx context.Context, name string) (proto int, err error) {
	return lookupProtocolMap(name)
}

func (*Resolver) lookupHost(ctx context.Context, host string) (addrs []string, err error) {
	return nil, syscall.ENOTSUP
}

func (*Resolver) lookupIP(ctx context.Context, network, host string) (addrs []IPAddr, err error) {
	return nil, syscall.ENOTSUP
}

func (*Resolver) lookupPort(ctx context.Context, network, service string) (port int, err error) {
	return 0, syscall.ENOTSUP
}

func (*Resolver) lookupCNAME(ctx context.Context, name string) (cname string, err error) {
	return "", syscall.ENOTSUP
}

func (*Resolver) lookupSRV(ctx context.Context, service, proto, name string) (cname string, srvs []*SRV, err error) {
	return "", nil, syscall.ENOTSUP
}

func (*Resolver) lookupMX(ctx context.Context, name string) (mxs []*MX, err error) {
	return nil, syscall.ENOTSUP
}

func (*Resolver) lookupNS(ctx context.Context, name string) (nss []*NS, err error) {
	return nil, syscall.ENOTSUP
}

func (*Resolver) lookupTXT(ctx context.Context, name string) (txts []string, err error) {
	return nil, syscall.ENOTSUP
}

func (*Resolver) lookupAddr(ctx context.Context, addr string) (ptrs []string, err error) {
	return nil, syscall.ENOTSUP
}

// concurrentThreadsLimit returns the number of threads we permit to
// run concurrently doing DNS lookups.
func concurrentThreadsLimit() int {
	return 1
}
