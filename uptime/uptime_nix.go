//go:build !windows
// +build !windows

package uptime

import (
	"syscall"
)

// Uptime returns the number of seconds since the system has booted
func Uptime() (int64, error) {
	info := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(info)
	if err != nil {
		return 0, err
	}
	return int64(info.Uptime), nil
}
