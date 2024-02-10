package dialer

import (
	"context"
	"fmt"
	"net"
	"time"
)

func DialTCP(hostname string, min, max int, invert bool) (OpenPorts []int) {
	for i := min; i <= max; i++ {
		connStr := fmt.Sprintf("%s:%d", hostname, i)

		conn, err := net.DialTimeout("tcp", connStr, time.Millisecond*225)
		if err != nil {
			if invert {
				OpenPorts = append(OpenPorts, i)
			}
			continue
		}
		defer conn.Close()

		if !invert {
			OpenPorts = append(OpenPorts, i)
		}

	}

	return
}

func DialUDP(hostname string, min, max int, invert bool) (OpenPorts []int) {
	for i := min; i <= max; i++ {
		ch := make(chan int)
		connStr := fmt.Sprintf("%s:%d", hostname, i)

		raddr, err := net.ResolveUDPAddr("udp", connStr)
		if err != nil {
			continue
		}

		conn, err := net.DialUDP("udp", nil, raddr)
		if err != nil {
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*225)
		defer cancel()
		go testUDP(conn, ch)

		select {
		case <-ctx.Done():
		case l := <-ch:
			if l > 0 && !invert {
				OpenPorts = append(OpenPorts, i)
			} else if l == 0 && invert {
				OpenPorts = append(OpenPorts, i)
			}
		}
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
