package config

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var Log *zap.SugaredLogger

func initLogger() {

	Log = CreateZapFactory(func(entry zapcore.Entry) error {
		return nil
	})
}

func CreateZapFactory(entry func(zapcore.Entry) error) *zap.SugaredLogger {

	// 以下才是 非调试（生产）模式所需要的代码
	encoderConfig := zap.NewProductionEncoderConfig()

	timePrecision := GlobalConfig.GetString("Logs.TimePrecision")
	var recordTimeFormat string
	switch timePrecision {
	case "second":
		recordTimeFormat = "2006-01-02 15:04:05"
	case "millisecond":
		recordTimeFormat = "2006-01-02 15:04:05.000"
	default:
		recordTimeFormat = "2006-01-02 15:04:05"

	}
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(recordTimeFormat))
	}
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.TimeKey = "created_at" // 生成json格式日志的时间键字段，默认为 ts,修改以后方便日志导入到 ELK 服务器

	var encoder zapcore.Encoder
	switch GlobalConfig.GetString("Logs.TextFormat") {
	case "console":
		encoder = zapcore.NewConsoleEncoder(encoderConfig) // 普通模式
	case "json":
		encoder = zapcore.NewJSONEncoder(encoderConfig) // json格式
	default:
		encoder = zapcore.NewConsoleEncoder(encoderConfig) // 普通模式
	}

	//写入器
	//fileName := variables.BasePath + config.GlobalConfig.GetString("Logs.DBMSLogName")
	var fileName string
	fileName = GlobalConfig.GetString("Logs.GinLogName")

	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,                               //日志文件的位置
		MaxSize:    GlobalConfig.GetInt("Logs.MaxSize"),    //在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: GlobalConfig.GetInt("Logs.MaxBackups"), //保留旧文件的最大个数
		MaxAge:     GlobalConfig.GetInt("Logs.MaxAge"),     //保留旧文件的最大天数
		Compress:   GlobalConfig.GetBool("Logs.Compress"),  //是否压缩/归档旧文件
	}
	writer := zapcore.AddSync(lumberJackLogger)
	// 开始初始化zap日志核心参数，
	//参数一：编码器
	//参数二：写入器
	//参数三：参数级别，debug级别支持后续调用的所有函数写日志，如果是 fatal 高级别，则级别>=fatal 才可以写日志
	zapCore := zapcore.NewCore(encoder, writer, zap.InfoLevel)
	return zap.New(zapCore, zap.AddCaller(), zap.Hooks(entry)).Sugar()
}
