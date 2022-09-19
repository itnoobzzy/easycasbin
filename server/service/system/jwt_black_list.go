package system

import (
	"akcasbin/server/global"
	"akcasbin/server/models"
	"context"
	"go.uber.org/zap"
	"time"
)

type JwtService struct{}

// JsonInBlackList 拉黑jwt
func (jwtService *JwtService) JsonInBlackList(jwtList models.JwtBlackList) (err error) {
	err = global.CASBIN_DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

// IsBlacklist 判断JWT是否在黑名单内部
func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
	// err := global.GVA_DB.Where("jwt = ?", jwt).First(&system.JwtBlacklist{}).Error
	// isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	// return !isNotFound
}

// GetRedisJWT 从redis取jwt
func (jwtService *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.CASBIN_REDIS.Get(context.Background(), userName).Result()
	return redisJWT, err
}

// SetRedisJWT jwt存入redis并设置过期时间
func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.CASBIN_CONFIG.JWT.ExpiresTime) * time.Second
	err = global.CASBIN_REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

func LoadAll() {
	var data []string
	err := global.CASBIN_DB.Model(&models.JwtBlackList{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.CASBIN_LOG.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	} // jwt黑名单 加入 BlackCache 中
}
