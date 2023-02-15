//go:build windows

package gos

import (
	"golang.org/x/sys/windows"
	"linktree_core/global"
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
	err := cmd.Start()
	if err != nil {
		global.GLOG.Panicf("The service background failed to start")
		return
	}
}
