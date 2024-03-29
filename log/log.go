package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// type Logger interface {
// 	Info(string, ...interface{})
// 	Fatal(string, ...interface{})
// 	Debug(string, ...interface{})
// 	Error(string, ...interface{})
// }

var logger *zap.Logger

func Init() {
	hook := lumberjack.Logger{
		Filename:   "./logs/api_server1.log", // 日志文件路径
		MaxSize:    128,                      // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 5,                        // 日志文件最多保存多少个备份
		MaxAge:     30,                       // 文件最多保存多少天
		Compress:   true,                     // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)
	// 新建一个ZapCore
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// 因为我们的代码对zap日志包提供的函数做了一层封装，所以需要在堆栈信息中skip一层
	skip := zap.AddCallerSkip(1)
	// 构造日志
	logger = zap.New(core, caller, skip, development, filed)
}

func SyncLogger() {
	logger.Sync()
}

// info level log
func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

// fatal level log
func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

// debug level log
func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

// error level log
func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

// warn level log
func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}
