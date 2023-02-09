package gos

import (
	"fmt"
	"github.com/duke-git/lancet/v2/random"
	"linktree_core/utils/dlog"
	"os"
	"strings"
)

// FirstPassFile Create a temporary password file
func FirstPassFile() {
	str := random.RandString(16)
	open, err := os.OpenFile("./pass", os.O_RDONLY, 0666)
	if err != nil {
		dlog.Log.Infof("Create a temporary password file")
		err1 := os.WriteFile("./pass", []byte(fmt.Sprintf("%s\n%s", str[0:8], str[8:16])), 0666)
		if err1 != nil {
			dlog.Log.Warn("Temporary password file creation failed")
			return
		}
		dlog.Log.Infof("username:\t%s", str[0:8])
		dlog.Log.Infof("password:\t%s", str[8:16])
		return
	}
	defer func(open *os.File) {
		err2 := open.Close()
		if err2 != nil {
			dlog.Log.Warn("Temporary password handle closed failed")
		}
	}(open)
}

// ReadPassFile 读取临时密码文件
func ReadPassFile() (error, string, string) {
	readFile, err1 := os.ReadFile("pass")
	if err1 != nil {
		return err1, "", ""
	}
	arr := strings.Fields(string(readFile))
	return nil, arr[0], arr[1]
}
