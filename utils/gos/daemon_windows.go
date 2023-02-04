//go:build windows

package gos

import (
	"golang.org/x/sys/windows"
	"os"
	"os/exec"
	"syscall"
)

func BackgroundStart() {
	cmd := &exec.Cmd{
		Path: os.Args[0],
		Args: []string{
			os.Args[0],
			"start",
		},
		SysProcAttr: &syscall.SysProcAttr{
			// windows环境的分离进程Flag
			CreationFlags: windows.DETACHED_PROCESS,
		},
	}
	cmd.Start()
}
