package interfacedialer

import (
	"net"
	"os"
)

func NewConn(fd int) (conn net.Conn, err error) {
	var (
		f = os.NewFile(uintptr(fd), "")
	)
	defer f.Close()

	if conn, err = net.FileConn(f); err != nil {
		return
	}

	return
}
