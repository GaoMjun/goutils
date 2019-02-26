package interfacedialer

import (
	"fmt"
	"log"
	"testing"
)

func TestDial(t *testing.T) {
	conn, err := Dial("tcp", "123.125.115.110:80", "en0", nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	fmt.Fprint(conn, "GET / HTTP/1.1\r\nHost: baidu.com\r\n\r\n")

	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(buffer[:n]))
}
