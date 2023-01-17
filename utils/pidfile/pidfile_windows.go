package pidfile

import (
	"os"
	"syscall"
)

func processExists(pid int) bool {
	proc, err := os.FindProcess(pid)
	if err != nil {
		return true
	}
	err = proc.Signal(syscall.SIGKILL) // 杀死进程
	if err != nil {
		return true
	}
	return false
}
