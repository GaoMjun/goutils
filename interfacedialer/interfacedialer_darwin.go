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
	"strconv"
	"strings"
	"syscall"
	"unsafe"
)

func Dial(network, address, ifname string, getProtectedSocket func(int, string, int) int) (conn *Conn, err error) {
	var (
		hostport = strings.Split(address, ":")
		host     = hostport[0]
		port, _  = strconv.Atoi(hostport[1])
		raddr    = syscall.SockaddrInet4{Port: port}
		fd       int
		ifname_c = C.CString(ifname)
		ifIdx    = int(C.if_nametoindex(ifname_c))
		laddr    = syscall.SockaddrInet4{}
	)

	C.free(unsafe.Pointer(ifname_c))
	copy(raddr.Addr[:], net.ParseIP(host).To4())

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

	err = syscall.SetsockoptInt(fd, syscall.IPPROTO_IP, syscall.IP_BOUND_IF, ifIdx)
	if err != nil {
		syscall.Close(fd)
		return
	}

	if err = syscall.Bind(fd, &laddr); err != nil {
		syscall.Close(fd)
		return
	}

	err = syscall.Connect(fd, &raddr)
	if err != nil {
		syscall.Close(fd)
		return
	}

	err = syscall.SetNonblock(fd, true)
	if err != nil {
		syscall.Close(fd)
		return
	}

	conn, err = NewConn(fd)
	return
}
