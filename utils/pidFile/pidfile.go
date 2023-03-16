// Package pidFile 创建pid文件，保证服务单列运行
package pidFile

import (
	"fmt"
	"os"
	"path/filepath"
)

// PIDFile is a file used to store the process ID of a running process.
type PIDFile struct {
	path string
}

// New 创建pid文件
func New(path string) (*PIDFile, error) {
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
