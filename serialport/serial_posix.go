// +build !windows

package serialport

import (
	"sort"
)

func getDefaultPort() string {
	ports := enumerateSerialPorts()
	if len(ports) == 0 {
		return ""
	}
	sort.Strings(ports)
	return ports[0]
}
