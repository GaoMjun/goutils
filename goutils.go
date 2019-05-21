package goutils

import (
	"errors"
	"fmt"
	"math/rand"
	"net"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func RepeatTimer(d time.Duration, f func()) {
	go func() {
		for range time.Tick(d) {
			f()
		}
	}()
}

func CreateFileNotExist(filename string) (f *os.File, err error) {
	var (
		base = path.Dir(filename)
	)

	_, err = os.Stat(base)
	if err != nil {
		err = os.MkdirAll(base, os.ModePerm)
		if err != nil {
			return
		}
	}

	_, err = os.Stat(filename)
	if err == nil {
		err = errors.New(fmt.Sprint("file exist ", filename))
		return
	}

	f, err = os.Create(filename)
	if err != nil {
		return
	}

	return
}

func RandBytes(length int) []byte {
	rand.Seed(time.Now().UnixNano())
	p := make([]byte, length)
	rand.Read(p)
	return p
}

func RandString(length int) string {
	return fmt.Sprintf("%x", RandBytes(length))
}

func SplitHostPort(hostport string) (host string, port int) {
	ss := strings.Split(hostport, ":")
	host = ss[0]

	if len(ss) == 2 {
		port, _ = strconv.Atoi(ss[1])
		return
	}

	port = 80
	return
}

func SplitIPPort(ipport string) (ip string, port int) {
	ss := strings.Split(ipport, ":")
	ip = ss[0]

	if len(ss) == 2 {
		port, _ = strconv.Atoi(ss[1])
		return
	}

	port = 0
	return
}

func InetNtoA(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

func InetAtoN(ip string) uint32 {
	ipBytes := net.ParseIP(ip).To4()
	return uint32(ipBytes[0])<<24 | uint32(ipBytes[1])<<16 | uint32(ipBytes[2])<<8 | uint32(ipBytes[3])
}

func InetNtoP(ip uint32) net.IP {
	return []byte{byte(ip & 0xFF000000 >> 24), byte(ip & 0x00FF0000 >> 16), byte(ip & 0x0000FF00 >> 8), byte(ip & 0x000000FF >> 0)}
}

func InetPtoA(ip net.IP) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ip[0]), byte(ip[1]), byte(ip[2]), byte(ip[3]))
}

func XOR(i, o, key []byte) {
	for i, b := range i {
		for j := 0; j < len(key); j++ {
			b ^= key[j]
		}

		o[i] = b
	}
	return
}
