package process

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

func KillProcess() error {
	pidByte, err := os.ReadFile("./pid")
	if err != nil {
		return errors.New("there is no pid file")
	}
	pidString := strings.TrimSpace(string(pidByte))
	pid, err1 := strconv.Atoi(pidString)
	if err1 != nil {
		return errors.New("PID conversion failed")
	}
	if !QuitProcess(pid) {
		return errors.New("kill fail")
	}
	return nil
}
