package system

import (
	"akcasbin/forms"
	"akcasbin/global"
	"akcasbin/models"
	"akcasbin/models/response"
	"akcasbin/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type UserApi struct{}

// Login 登录
func (ua *UserApi) Login(c *gin.Context) {
	var loginForm forms.Login
	if err := c.ShouldBindJSON(&loginForm); err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	if err := utils.Verify(loginForm, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	u := &models.User{Username: loginForm.Username, Password: loginForm.Password}
	if user, err := userService.Login(u); err != nil {
		global.CASBIN_LOG.Error("login failed! username or password wrong!", zap.Error(err))
		response.FailWithMessage("用户名不存在或者密码错误", c)
	} else {
		if user.Active != 1 {
			global.CASBIN_LOG.Error("login failed! not active!")
			response.FailWithMessage("用户账号未被激活!", c)
			return
		}
		ua.TokenNext(c, *user)
	}
}

// TokenNext 登录后签发JWT
func (ua *UserApi) TokenNext(c *gin.Context, user models.User) {
	j := &utils.JWT{SigningKey: []byte(global.CASBIN_CONFIG.JWT.SigningKey)}
	claims := j.CreateClaims(models.BaseClaims{
		ID:       user.ID,
		Username: user.Username,
		NickName: user.AliasName,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.CASBIN_LOG.Error("get token failed!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	// 如果是单点登录，无需redis存储token
	if !global.CASBIN_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "login success!", c)
		return
	}
	if jwt, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			global.CASBIN_LOG.Error("set login state failed!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败！", c)
			return
		}
		response.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "login success!", c)
	} else if err != nil {
		global.CASBIN_LOG.Error("set login state failed!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败！", c)
	} else {
		// 废弃旧JWT
		var blackJWT models.JwtBlackList
		blackJWT.Jwt = jwt
		if err := jwtService.JsonInBlackList(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "login success!", c)
	}
}

// 注册
func (ua *UserApi) Register(c *gin.Context) {
	var registerForm forms.Register
	if err := c.ShouldBindJSON(&registerForm); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user := &models.User{
		Username:  registerForm.Username,
		Password:  registerForm.Password,
		AliasName: registerForm.NickName,
	}
	res, err := userService.Register(*user)
	if err != nil {
		global.CASBIN_LOG.Error("register failed!", zap.Error(err))
		response.FailWithDetailed(response.RegisterResponse{res}, "register failed!", c)
	} else {
		response.OkWithDetailed(response.RegisterResponse{res}, "register success!", c)
	}
}
