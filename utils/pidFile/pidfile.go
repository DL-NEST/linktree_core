// Package pidFile 创建pid文件，保证服务单列运行
package pidFile

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// PIDFile is a file used to store the process ID of a running process.
type PIDFile struct {
	path string
}

func checkPIDFileAlreadyExists(path string) error {
	if pidByte, err := os.ReadFile(path); err == nil {
		pidString := strings.TrimSpace(string(pidByte))
		if pid, err1 := strconv.Atoi(pidString); err1 == nil {

			if ProcessExists(pid) {
				return fmt.Errorf("pid file found %s", path)
			}

		}
	}
	return nil
}

// New 创建pid文件
func New(path string) (*PIDFile, error) {
	if err := checkPIDFileAlreadyExists(path); err != nil {
		panic(err)
	}

	if err := os.MkdirAll(filepath.Dir(path), os.FileMode(0755)); err != nil {
		return nil, err
	}
	if err := os.WriteFile(path, []byte(fmt.Sprintf("%d", os.Getpid())), 0644); err != nil {
		return nil, err
	}
	return &PIDFile{path: path}, nil
}

// Remove removes the PIDFile.
func (file PIDFile) Remove() error {
	return os.Remove(file.path)
}

func ProcessExists(pid int) bool {
	proc, err := os.FindProcess(pid)
	// 进程不存在
	if err != nil {
		return false
	}
	err = proc.Kill()
	if err != nil {
		return true
	}
	return false
}

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
	if !ProcessExists(pid) {
		return errors.New("kill fail")
	}
	return nil
}
