package logger

import (
	"sync"

	"go.uber.org/zap"
)

var (
	once     sync.Once
	instance *zap.Logger
)

// GetLogger 返回单例 Logger
func GetLogger() *zap.Logger {
	once.Do(func() {
		var err error
		instance, err = zap.NewProduction() // 初始化 Logger
		if err != nil {
			panic(err) // 初始化失败时 panic
		}

		// 添加全局字段
		instance = instance.With(
			zap.String("app", "nebula"),
		)
	})

	return instance
}

// Sync 刷新日志缓冲区
func Sync() {
	if instance != nil {
		_ = instance.Sync()
	}
}
