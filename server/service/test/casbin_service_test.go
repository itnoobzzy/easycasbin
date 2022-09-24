package test

import (
	"akcasbin/core"
	"akcasbin/global"
	"akcasbin/initialize"
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"testing"

	"github.com/casbin/casbin/v2/model"
)

func setup() {
	config := "../../config.yaml"
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	if err = v.Unmarshal(&global.CASBIN_CONFIG); err != nil {
		fmt.Println(err)
	}
	global.CASBIN_LOG = core.Zap()
	zap.ReplaceGlobals(global.CASBIN_LOG)
	global.CASBIN_DB = initialize.Gorm()
}

func TestCasbinService(t *testing.T) {
	setup()
	a, err := gormadapter.NewAdapterByDB(global.CASBIN_DB)
	if err != nil {
		t.Errorf("error: adapter: %s", err)
	}
	m, err := model.NewModelFromFile("../../casbin_rbac_domain.conf")
	if err != nil {
		t.Errorf("error: model: %s", err)
	}
	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		t.Errorf("error: enforcer: %s", err)
	}

	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"
	_, err = e.Enforce(sub, obj, act)
}
