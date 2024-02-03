package main

import (
	"flag"
	"fmt"
	"os"
	"pp-scanner/dialer"
)

var (
	hostname string
	minport  int
	maxport  int
)

func main() {
	flag.StringVar(&hostname, "h", "", "-h [hostname]")
	flag.IntVar(&minport, "min", 1, "-min [lowest port #]")
	flag.IntVar(&maxport, "max", 65535, "-max [highest port #]")
	flag.Parse()

	if hostname == "" {
		fmt.Fprintln(os.Stderr, "Error: no hostname specified")
		os.Exit(1)
	}

	if minport > maxport {
		fmt.Fprintln(os.Stderr, "Error: minport greater than maxport")
		os.Exit(1)
	}

	ports := dialer.Dial(hostname, minport, maxport)
	fmt.Println(ports)
}
