package core

import (
	"akcasbin/core/internal"
	"akcasbin/global"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"os"
	"time"
)

// Viper //
// 优先级: 命令行 > 环境变量 > 默认值
func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" {
			// 当命令行参数为空的时候判断环境变量是否有CASBIN_CONFIG
			// 如果有就将环境变量值作为config
			// 否则根据gin 启动模式读取对应的配置文件
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDefaultFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigDefaultFile)
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigReleaseFile)
				case gin.TestMode:
					config = internal.ConfigTestFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigTestFile)
				}
			} else {
				config = configEnv
				fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", internal.ConfigEnv, config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", config)
		}

	} else {
		// 将函数传递的第一个参数作为config
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.CASBIN_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.CASBIN_CONFIG); err != nil {
		fmt.Println(err)
	}
	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(time.Second * time.Duration(global.CASBIN_CONFIG.JWT.ExpiresTime)),
	)
	return v
}
