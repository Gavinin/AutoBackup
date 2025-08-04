package log

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var Logger *zap.SugaredLogger

func Init() {
	fileWriter, _ := rotatelogs.New(
		"logs/app-%Y-%m-%d.log",
		rotatelogs.WithLinkName("logs/app.log"),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	// 2. 控制台输出
	consoleWriter := zapcore.Lock(os.Stdout)

	// 3. 编码器
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(fileWriter),
		zap.InfoLevel,
	)

	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		consoleWriter,
		zap.InfoLevel,
	)

	// 4. 合并两个 core
	core := zapcore.NewTee(fileCore, consoleCore)

	logger := zap.New(core)
	defer logger.Sync()
	Logger = logger.Sugar()
}
