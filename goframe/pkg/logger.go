package pkg

import (
	"fmt"
	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"goframe-code/config"
	"goframe-code/global"
	"os"
	"path"
	"time"
)

var level zapcore.Level

func init() {

	fmt.Println("初始化Log")

	global.Log = Zap()
}

func Zap() (logger *zap.Logger) {
	switch config.GLOBAL_CONF.Log.Level { // 初始化配置文件的Level
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(getEncoderCore(), zap.AddStacktrace(level))
	} else {
		logger = zap.New(getEncoderCore())
	}
	if config.GLOBAL_CONF.Log.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (c zapcore.EncoderConfig) {
	c = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  config.GLOBAL_CONF.Log.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case config.GLOBAL_CONF.Log.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		c.EncodeLevel = zapcore.LowercaseLevelEncoder
	case config.GLOBAL_CONF.Log.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		c.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case config.GLOBAL_CONF.Log.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		c.EncodeLevel = zapcore.CapitalLevelEncoder
	case config.GLOBAL_CONF.Log.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		c.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		c.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return c
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if config.GLOBAL_CONF.Log.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore() (core zapcore.Core) {
	writer, err := GetWriteSyncer() // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	return zapcore.NewCore(getEncoder(), writer, level)
}

// CustomTimeEncoder 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(config.GLOBAL_CONF.Log.Prefix + "2006/01/02 - 15:04:05.000"))
}

func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		path.Join(config.GLOBAL_CONF.Log.Director, "%Y-%m-%d.log"),
		zaprotatelogs.WithLinkName(config.GLOBAL_CONF.Log.LinkName),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)
	if config.GLOBAL_CONF.Log.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
