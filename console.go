package licheejs

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/dop251/goja"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ConsoleCallBack = func(args ...any)

type console struct {
	logger   *zap.Logger
	CallBack ConsoleCallBack
}

type LogOutMode int

const (
	LOM_RELEASE LogOutMode = iota
	LOM_DEBUG
)

func getJsonEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 修改时间编码器

	// 在日志文件中使用大写字母记录日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// NewConsoleEncoder 打印更符合人们观察的方式
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getConsoleEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 修改时间编码器

	// 在日志文件中使用大写字母记录日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// NewConsoleEncoder 打印更符合人们观察的方式
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(path string) (zapcore.WriteSyncer, func(), error) {
	return zap.Open(path)
}

func newZap(path string, mod LogOutMode) (*zap.Logger, func(), error) {
	ws, close, err := getLogWriter(path)
	if err != nil {
		return nil, nil, err
	}

	var core zapcore.Core

	if mod == LOM_DEBUG {
		core = zapcore.NewTee(
			zapcore.NewCore(getConsoleEncoder(), zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
			zapcore.NewCore(getJsonEncoder(), ws, zap.DebugLevel),
		)
	} else {
		core = zapcore.NewTee(
			zapcore.NewCore(getJsonEncoder(), ws, zap.DebugLevel),
		)
	}

	log := zap.New(core, zap.AddCaller())
	return log, close, nil
}

func newConsole(log *zap.Logger, cb ConsoleCallBack) *console {
	c := &console{logger: log}
	c.CallBack = cb
	return c
}

func (c console) log(level zapcore.Level, args ...goja.Value) {
	var strs strings.Builder
	for i := 0; i < len(args); i++ {
		if i > 0 {
			strs.WriteString(" ")
		}
		strs.WriteString(c.valueString(args[i]))
	}
	msg := strs.String()

	flag := ""

	switch level { //nolint:exhaustive
	case zapcore.DebugLevel:
		{
			c.logger.Debug(msg)
			flag = "[debug]"
		}
	case zapcore.InfoLevel:
		{
			c.logger.Info(msg)
			flag = "[info]"
		}
	case zapcore.WarnLevel:
		{
			c.logger.Warn(msg)
			flag = "[warn]"
		}
	case zapcore.ErrorLevel:
		{
			c.logger.Error(msg)
			flag = "[error]"
		}
	}

	// 写入回调
	if c.CallBack != nil {
		c.CallBack(fmt.Sprintf("%s %s", flag, msg))
	}
}

func (c console) Log(args ...goja.Value) {
	c.Info(args...)
}

func (c console) Debug(args ...goja.Value) {
	c.log(zapcore.DebugLevel, args...)
}

func (c console) Info(args ...goja.Value) {
	c.log(zapcore.InfoLevel, args...)
}

func (c console) Warn(args ...goja.Value) {
	c.log(zapcore.WarnLevel, args...)
}

func (c console) Error(args ...goja.Value) {
	c.log(zapcore.ErrorLevel, args...)
}

func (c console) valueString(v goja.Value) string {
	mv, ok := v.(json.Marshaler)
	if !ok {
		return v.String()
	}

	b, err := json.Marshal(mv)
	if err != nil {
		return v.String()
	}
	return string(b)
}
