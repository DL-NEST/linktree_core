package glog

import (
	"github.com/fatih/color"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _levelToColor = map[zapcore.Level]func(a ...interface{}) string{
	zap.DebugLevel:  color.New(color.FgHiMagenta).Add(color.Bold).SprintFunc(),
	zap.InfoLevel:   color.New(color.FgCyan).Add(color.Bold).SprintFunc(),
	zap.WarnLevel:   color.New(color.FgYellow).Add(color.Bold).SprintFunc(),
	zap.ErrorLevel:  color.New(color.FgRed).Add(color.Bold).SprintFunc(),
	zap.DPanicLevel: color.New(color.BgRed).Add(color.Bold).SprintFunc(),
	zap.PanicLevel:  color.New(color.BgRed).Add(color.Bold).SprintFunc(),
	zap.FatalLevel:  color.New(color.BgRed).Add(color.Bold).SprintFunc(),
}
