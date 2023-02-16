//go:build darwin
// +build darwin

package interfacedialer

/*
#include <net/if.h>
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"fmt"
	"net"
	"syscall"
)

func Dial(network string, ip net.IP, port int, ifname string, ifIdx int, getProtectedSocket func(int, string, int) int) (conn net.Conn, err error) {
	var (
		raddr = syscall.SockaddrInet4{Port: port}
		fd    int
	)

	copy(raddr.Addr[:], ip)

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

	if err = syscall.SetsockoptInt(fd, syscall.IPPROTO_IP, syscall.IP_BOUND_IF, ifIdx); err != nil {
		syscall.Close(fd)
		return
	}

	if err = syscall.Bind(fd, &syscall.SockaddrInet4{}); err != nil {
		syscall.Close(fd)
		return
	}

	if err = syscall.Connect(fd, &raddr); err != nil {
		syscall.Close(fd)
		return
	}

	if err = syscall.SetNonblock(fd, true); err != nil {
		syscall.Close(fd)
		return
	}

	conn, err = NewConn(fd)
	return
}
