package main

import (
	"flag"
	"fmt"
	"os"
	"pp-scanner/dialer"
	"pp-scanner/formatoutput"
)

var (
	hostname string
	minport  int
	maxport  int
	udp      bool
)

func main() {
	flag.StringVar(&hostname, "h", "", "-h [hostname]")
	flag.IntVar(&minport, "min", 1, "-min [lowest port #]")
	flag.IntVar(&maxport, "max", 65535, "-max [highest port #]")
	flag.BoolVar(&udp, "u", false, "-u [scan for udp portss]")
	flag.Parse()

	if hostname == "" {
		fmt.Fprintln(os.Stderr, "Error: no hostname specified")
		os.Exit(1)
	}
	if minport > maxport {
		fmt.Fprintln(os.Stderr, "Error: minport greater than maxport")
		os.Exit(1)
	}

	var ports []int
	if !udp {
		ports = dialer.DialTCP(hostname, minport, maxport)
	} else {
		ports = dialer.DialUDP(hostname, minport, maxport)
	}
	fmt.Printf("%s", formatoutput.FormatOutput(ports))
}
