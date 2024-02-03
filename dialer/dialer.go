package dialer

import (
	"net"
	"strconv"
	"time"
)

func Dial(hostname string, min, max int) (OpenPorts []int) {
	for i := min; i <= max; i++ {
		connStr := hostname + ":" + strconv.Itoa(i)

		conn, err := net.DialTimeout("tcp", connStr, time.Second)
		if err != nil {
			continue
		}
		defer conn.Close()
		OpenPorts = append(OpenPorts, i)
	}

	return
}
