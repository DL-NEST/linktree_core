//go:build linux && darwin

package process

import (
	"os"
	"syscall"
)

func QuitProcess(pid int) bool {
	proc, err := os.FindProcess(pid)
	// unix中存在不存在都会返回
	if err != nil {
		return false
	}
	err = proc.Signal(syscall.SIGINT)
	if err != nil {
		return true
	}
	return false
}
