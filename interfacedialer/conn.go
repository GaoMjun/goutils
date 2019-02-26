package interfacedialer

import (
	"net"
	"os"
)

type Conn struct {
	net.Conn
}

func NewConn(fd int) (conn *Conn, err error) {
	f := os.NewFile(uintptr(fd), "")
	c, err := net.FileConn(f)
	f.Close()
	if err != nil {
		return
	}

	conn = &Conn{c}
	return
}
