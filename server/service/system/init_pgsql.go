package system

import (
	"akcasbin/server/config"
	"akcasbin/server/forms"
	"akcasbin/server/global"
	"akcasbin/server/utils"
	"context"
	"errors"
	"fmt"
	"github.com/gookit/color"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgsqlInitHandler struct{}

func NewPgsqlInitHandler() *PgsqlInitHandler {
	return &PgsqlInitHandler{}
}

// WriteConfig pgsql 回写配置
func (h PgsqlInitHandler) WriteConfig(ctx context.Context) error {
	c, ok := ctx.Value("config").(config.Pgsql)
	if !ok {
		return errors.New("postgresql config invalid")
	}
	global.CASBIN_CONFIG.System.DbType = "pgsql"
	global.CASBIN_CONFIG.Pgsql = c
	global.CASBIN_CONFIG.JWT.SigningKey = uuid.NewV4().String()
	cs := utils.StructToMap(global.CASBIN_CONFIG)
	for k, v := range cs {
		global.CASBIN_VP.Set(k, v)
	}
	return global.CASBIN_VP.WriteConfig()
}

// EnsureDB 创建数据库并初始化 pg
func (h PgsqlInitHandler) EnsureDB(ctx context.Context, conf *forms.InitDB) (next context.Context, err error) {
	if s, ok := ctx.Value("dbtype").(string); !ok || s != "pgsql" {
		return ctx, ErrDBTypeMismatch
	}
	c := conf.ToPgsqlConfig()
	next = context.WithValue(ctx, "config", c)
	if c.Dbname == "" {
		return ctx, nil
	}
	// 创建数据库
	dsn := conf.PgsqlEmptyDsn()
	createSql := fmt.Sprintf("CREATE DATABASE %s;", c.Dbname)
	if err = createDataBase(dsn, "pgx", createSql); err != nil {
		return nil, err
	}
	// 创建链接session
	var db *gorm.DB
	if db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  c.Dsn(),
		PreferSimpleProtocol: false,
	}), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return ctx, err
	}
	next = context.WithValue(next, "db", db)
	return next, err
}

func (h PgsqlInitHandler) InitTables(ctx context.Context, inits initSlice) error {
	return createTables(ctx, inits)
}

func (h PgsqlInitHandler) InitData(ctx context.Context, inits initSlice) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) { c() }(cancel)
	for i := 0; i < len(inits); i++ {
		if inits[i].DataInserted(next) {
			color.Info.Printf(InitDataExist, Pgsql, inits[i].InitializerName())
			continue
		}
		if n, err := inits[i].InitializeData(next); err != nil {
			color.Info.Printf(InitDataFailed, Pgsql, inits[i].InitializerName())
			return err
		} else {
			next = n
			color.Info.Printf(InitDataSuccess, Pgsql, inits[i].InitializerName())
		}
	}
	color.Info.Printf(InitSuccess, Pgsql)
	return nil
}
