package logger

import (
	"github.com/fatih/color"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"linktree_core/commands"
	"os"
	"time"
)

var Log *zap.SugaredLogger

func init() {
	SetupLogger()
}

func SetupLogger() {
	stdout := zapcore.NewCore(getStdoutEncoder(), zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	saveLog := zapcore.NewCore(getSaveEncoder(), zapcore.AddSync(getLogWriter()), zapcore.DebugLevel)
	logger := zap.New(zapcore.NewTee(stdout, saveLog), zap.AddCaller())
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	Log = logger.Sugar()
}

func getSaveEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.NanosDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getStdoutEncoder() zapcore.Encoder {
	// 自定义时间输出格式
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		s := color.New(color.FgHiBlack).Add(color.Bold).Sprintf("[" + t.Format("2006-01-02 15:04:05") + "]")
		enc.AppendString(s)
	}
	// 自定义日志级别显示
	customLevelEncoder := func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		s := _levelToColor[level]("[" + level.CapitalString() + "]")
		enc.AppendString(s)
	}
	// 自定义文件：行号输出项
	customCallerEncoder := func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(caller.TrimmedPath())
	}
	encoderConf := zapcore.EncoderConfig{
		CallerKey:      "caller_line", // 打印文件名和行数
		LevelKey:       "level_name",
		MessageKey:     "msg",
		TimeKey:        "ts",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     customTimeEncoder,   // 自定义时间格式
		EncodeLevel:    customLevelEncoder,  // 小写编码器
		EncodeCaller:   customCallerEncoder, // 全路径编码器
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	return zapcore.NewConsoleEncoder(encoderConf)
}

func getLogWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   commands.LogPath + "/runtime.log",
		MaxSize:    2,     //每个文件最大占用内存，单位M
		MaxBackups: 5,     //备份数量
		MaxAge:     30,    //备份的天数
		Compress:   false, //是否压缩
		LocalTime:  true,  // 使用本地时间
	})
}
