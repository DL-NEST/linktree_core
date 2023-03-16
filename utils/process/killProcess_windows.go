//go:build windows

package process

import (
	"os"
)

func QuitProcess(pid int) bool {
	proc, err := os.FindProcess(pid)
	// windows中可以判断进程是否存在（进程结束了但是还是会判断true）,unix中存在不存在都会返回
	if err != nil {
		return false
	}
	err = proc.Kill()
	if err != nil {
		return true
	}
	return false
}
