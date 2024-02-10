package formatoutput

import "fmt"

func FormatOutput(arr []int, invert bool) string {
	var formatted string
	for _, val := range arr {
		if invert {
			formatted += fmt.Sprintf("Port %d status: closed\n", val)
		} else {
			formatted += fmt.Sprintf("Port %d status: open\n", val)
		}
	}

	return formatted
}
