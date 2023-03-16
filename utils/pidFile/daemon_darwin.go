//go:build darwin

package pidFile

import (
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
			Setsid: true,
		},
	}
	cmd.Start()
}
