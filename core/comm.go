package core

import "syscall"

type FDConn struct {
	Fd int
}

func (f FDConn) Write(b []byte) (int, error) {
	return syscall.Write(f.Fd, b)
}

func (f FDConn) Read(b []byte) (int, error) {
	return syscall.Read(f.Fd, b)
}