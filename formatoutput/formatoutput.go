package formatoutput

import "fmt"

func FormatOutput(arr []int) string {
	var formatted string
	for _, val := range arr {
		formatted += fmt.Sprintf("Port %d status: open\n", val)
	}

	return formatted
}
