// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package outside

import (
	"syscall"
)

type sproc syscall.Proc

func load(n string) (sd *sdll, e error) {
	d, e := syscall.LoadDLL(n)
	sd = (*sdll)(d)
	return
}

type sdll syscall.DLL

func (sd *sdll) mustFindProc(s string) *sproc {
	return (*sproc)((*syscall.DLL)(sd).MustFindProc(s))
}

//TODO(t): add error handling
func (sd *sdll) release() {
	(*syscall.DLL)(sd).Release()
}

func (sd *sdll) findProc(n string) (sp *sproc, e error) {
	p, e := (*syscall.DLL)(sd).FindProc(n)
	sp = (*sproc)(p)
	return
}

func newCallback(cb interface{}) uintptr {
	return syscall.NewCallback(cb)
}

func (sp *sproc) addr() uintptr { return (*syscall.Proc)(sp).Addr() }
