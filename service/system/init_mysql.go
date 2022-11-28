package system

import "C"
import (
	"easycasbin/config"
	"easycasbin/forms"
	"easycasbin/global"
	"easycasbin/utils"
	"context"
	"errors"
	"fmt"
	"github.com/gookit/color"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlInitHandler struct{}

func NewMysqlInitHandler() *MysqlInitHandler {
	return &MysqlInitHandler{}
}

// WriteConfig mysql 回写配置
func (h MysqlInitHandler) WriteConfig(ctx context.Context) error {
	c, ok := ctx.Value("config").(config.Mysql)
	if !ok {
		return errors.New("mysql config invalid")
	}
	global.CASBIN_CONFIG.System.DbType = "mysql"
	global.CASBIN_CONFIG.Mysql = c
	global.CASBIN_CONFIG.JWT.SigningKey = uuid.NewV4().String()
	cs := utils.StructToMap(global.CASBIN_CONFIG)
	for k, v := range cs {
		global.CASBIN_VP.Set(k, v)
	}
	return global.CASBIN_VP.WriteConfig()
}

// EnsureDB 创建数据库并初始化 mysql
func (h MysqlInitHandler) EnsureDB(ctx context.Context, conf *forms.InitDB) (next context.Context, err error) {
	if s, ok := ctx.Value("dbtype").(string); !ok || s != "mysql" {
		return ctx, ErrDBTypeMismatch
	}
	c := conf.ToMysqlConfig()
	next = context.WithValue(ctx, "config", c)
	// 如果没有数据库名，跳出初始化数据
	if c.Dbname == "" {
		return ctx, nil
	}

	// 创建数据库
	dsn := conf.MysqlEmptyDsn()
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", c.Dbname)
	if err = createDataBase(dsn, "mysql", createSql); err != nil {
		return nil, err
	}

	var db *gorm.DB
	if db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       c.Dsn(),
		DefaultStringSize:         191,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return ctx, err
	}
	next = context.WithValue(next, "db", db)
	return next, err
}

func (h MysqlInitHandler) InitTables(ctx context.Context, inits initSlice) error {
	return createTables(ctx, inits)
}

func (h MysqlInitHandler) InitData(ctx context.Context, inits initSlice) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) { c() }(cancel)
	for _, init := range inits {
		if init.DataInserted(next) {
			color.Info.Printf(InitDataExist, Mysql, init.InitializerName())
			continue
		}
		if n, err := init.InitializeData(next); err != nil {
			color.Info.Printf(InitDataFailed, Mysql, init.InitializerName())
			return err
		} else {
			next = n
			color.Info.Printf(InitDataSuccess, Mysql, init.InitializerName())
		}
	}
	color.Info.Printf(InitSuccess, Mysql)
	return nil
}
