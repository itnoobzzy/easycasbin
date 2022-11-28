package core

import (
	"easycasbin/core/internal"
	"easycasbin/global"
	"easycasbin/utils"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Zap 获取 zap.Logger
func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.CASBIN_CONFIG.Zap.Director); !ok {
		fmt.Printf("create %v directory\n", global.CASBIN_CONFIG.Zap.Director)
		err := os.Mkdir(global.CASBIN_CONFIG.Zap.Director, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.CASBIN_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
