package dialer

import (
	"context"
	"fmt"
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
		case l := <-ch:
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
	buf := make([]byte, 8192)

	_, err := conn.Read(buf)
	if err != nil {
		ch <- 0
		return
	}
	ch <- len(buf)
}
