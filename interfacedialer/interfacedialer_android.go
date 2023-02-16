//go:build android
// +build android

package interfacedialer

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"syscall"

	"github.com/GaoMjun/goutils"
)

func Dial(network, address, ifname string, getProtectedSocket func(int, string, int) int) (conn net.Conn, err error) {
	var (
		hostport = strings.Split(address, ":")
		host     = hostport[0]
		port, _  = strconv.Atoi(hostport[1])
		addr     syscall.SockaddrInet4
		fd       int
	)

	copy(addr.Addr[:], goutils.InetNtoP(goutils.InetAtoN(host)))
	addr.Port = port

	if network == "tcp" {
		fd = getProtectedSocket(0, host, port)
	} else if network == "udp" {
		fd = getProtectedSocket(1, host, port)
	} else {
		err = errors.New(fmt.Sprint("not support protocol", network))
	}
	if fd <= 0 {
		err = errors.New("getProtectedSocket failed")
	}
	if err != nil {
		return
	}

	conn, err = NewConn(fd)
	return
}
