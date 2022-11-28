package internal

import (
	"easycasbin/global"
	"fmt"
	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch global.CASBIN_CONFIG.System.DbType {
	case "mysql":
		logZap = global.CASBIN_CONFIG.Mysql.LogZap
	case "pgsql":
		logZap = global.CASBIN_CONFIG.Pgsql.LogZap
	}
	if logZap {
		global.CASBIN_LOG.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
