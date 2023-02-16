//go:build linux
// +build linux

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

func Dial(network, address, ifname string, getProtectedSocket func(int, string, int) int) (conn *net.Conn, err error) {
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
		fd, err = syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_IP)
	} else if network == "udp" {
		fd, err = syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_IP)
	} else {
		err = errors.New(fmt.Sprint("not support protocol", network))
	}
	if err != nil {
		return
	}

	err = syscall.SetsockoptString(fd, syscall.SOL_SOCKET, syscall.SO_BINDTODEVICE, ifname)
	if err != nil {
		syscall.Close(fd)
		return
	}

	err = syscall.Connect(fd, &addr)
	if err != nil {
		syscall.Close(fd)
		return
	}

	conn, err = NewConn(fd)
	return
}
