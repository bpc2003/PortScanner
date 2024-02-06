package dialer

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"strconv"
	"time"
)

func DialTCP(hostname string, min, max int) (OpenPorts []int) {
	for i := min; i <= max; i++ {
		connStr := hostname + ":" + strconv.Itoa(i)

		conn, err := net.DialTimeout("tcp", connStr, time.Millisecond*225)
		if err != nil {
			continue
		}
		defer conn.Close()
		OpenPorts = append(OpenPorts, i)
	}

	return
}

func DialUDP(hostname string, min, max int) (OpenPorts []int) {
	for i := min; i <= max; {
		fmt.Println(i)

		ch := make(chan int)
		connStr := hostname + ":" + strconv.Itoa(i)

		raddr, err := net.ResolveUDPAddr("udp", connStr)
		if err != nil {
			continue
		}

		conn, err := net.DialUDP("udp", nil, raddr)
		if err != nil {
			continue
		}
		defer conn.Close()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		go testUDP(conn, ch)

		select {
		case <-ctx.Done():
			fmt.Println("op cancelled")
		case l := <-ch:
			fmt.Println(l)
			if l > 0 {
				OpenPorts = append(OpenPorts, i)
			}
		}
		i++
	}

	return
}

func testUDP(conn *net.UDPConn, ch chan int) {
	fmt.Fprintln(conn, "PING")

	var buf bytes.Buffer
	io.Copy(&buf, conn)

	ch <- buf.Len()
}
