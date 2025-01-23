package main

import (
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)
type t_logger struct{
    logDir string
}

func Logger(logDir string) *t_logger{
    return &t_logger{
        logDir: logDir,
    }
}

func(t *t_logger) Create(logName string ,level logrus.Level) *logrus.Logger {
	logger := logrus.New()
	// 设置日志级别
	logger.SetLevel(logrus.InfoLevel)
	// 设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

    filePath := filepath.Join(t.logDir, logName+".log")
	// 设置日志输出到文件并进行切割
	logger.SetOutput(&lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    20, // 每个日志文件最大 20MB
		MaxBackups: 14, // 保留 14 个旧日志文件
		MaxAge:     7,  // 保留 7 天的日志文件
		Compress:   true,
	})

	return logger
}

// func main() {
// 	// 初始化日志记录器
// 	logger := Logger("./logs")
//     test := logger.Create("test",logrus.InfoLevel)
//     test.Debug("debug")
//     test.Info("info")
//     test.Warn("warn")
//     test.Error("error")
//     test.Fatal("fatal")
// }