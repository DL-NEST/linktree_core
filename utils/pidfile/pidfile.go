// Package pidfile provides structure and helper functions to create and remove
// PID file. A PID file is usually a file used to store the process ID of a
// running process.
package pidfile

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

// issues:用户手动修改pid文件程序就会报错
func checkPIDFileAlreadyExists(path string) error {
	if pidByte, err := os.ReadFile(path); err == nil {
		pidString := strings.TrimSpace(string(pidByte))
		if pid, err1 := strconv.Atoi(pidString); err1 == nil {
			if ProcessExists(pid) {
				return fmt.Errorf("pid file found, ensure gmqtt is not running or delete %s", path)
			}
		}
	}
	return nil
}

// New creates a PIDfile using the specified path.
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
