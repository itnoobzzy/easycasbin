package internal

import (
	"akcasbin/global"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var Zap = new(_zap)

type _zap struct{}

// GetEncoder 获取 zapcore.Encoder
func (z *_zap) GetEncoder() zapcore.Encoder {
	if global.CASBIN_CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(z.GetEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(z.GetEncoderConfig())
}

// GetEncoderConfig 获取zapcore.EncoderConfig
func (z *_zap) GetEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.CASBIN_CONFIG.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    global.CASBIN_CONFIG.Zap.ZapEncodeLevel(),
		EncodeTime:     z.CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
}

// CustomTimeEncoder 自定义日志输出时间格式
func (z *_zap) CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(global.CASBIN_CONFIG.Zap.Prefix + t.Format("2006/01/02 - 15:04:05.000"))
}

// GetEncoderCore 获取Encoder的 zapcore.Core
func (z *_zap) GetEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	writer, err := FileRotatelogs.GetWriteSyncer(l.String())
	if err != nil {
		fmt.Printf("Get Write Sycner Failed err:%v", err.Error())
		return nil
	}
	return zapcore.NewCore(z.GetEncoder(), writer, level)
}

// GetZapCores 根据配置文件的Level获取 []zapcore.Core
func (z *_zap) GetZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := global.CASBIN_CONFIG.Zap.TransportLevel(); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, z.GetEncoderCore(level, z.GetLevelPriority(level)))
	}
	return cores
}

// GetLevelPriority 根据 zapcore.Level 获取 zap.LevelEnablerFunc
func (z *_zap) GetLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool {
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool {
			return level == zapcore.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool {
			return level == zapcore.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool {
			return level == zapcore.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool {
			return level == zapcore.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool {
			return level == zapcore.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool {
			return level == zapcore.DebugLevel
		}
	default:
		return func(level zapcore.Level) bool {
			return level == zap.DebugLevel
		}
	}
}
