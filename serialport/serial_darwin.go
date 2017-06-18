package serialport

import (
	"path/filepath"
	"strings"
)

func enumerateSerialPorts() []string {
	list, _ := filepath.Glob("/dev/cu.*")
	var filteredList []string
	for _, s := range list {
		if !strings.Contains(s, "Bluetooth-") && !strings.Contains(s, "-Wireless") {
			filteredList = append(filteredList, s)
		}
	}
	return filteredList
}
