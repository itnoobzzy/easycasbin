package system

import (
	"easycasbin/global"
	"easycasbin/models"
	"easycasbin/models/request"
	"easycasbin/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// @function: Register
// @description: 用户注册
// @param: u model.SysUser
// @return: userInter system.SysUser, err error

type UserService struct{}

// Register Register
func (UserService *UserService) Register(u models.User) (userInfo models.User, err error) {
	var user models.User
	if !errors.Is(global.CASBIN_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInfo, errors.New("用户名已注册")
	}
	u.Password = utils.BcryptHash(u.Password)
	u.Active = 0
	err = global.CASBIN_DB.Create(&u).Error
	return u, err
}

// Login login
func (UserService *UserService) Login(u *models.User) (userInfo *models.User, err error) {
	if nil == global.CASBIN_DB {
		return nil, fmt.Errorf("db not init")
	}
	var user models.User
	err = global.CASBIN_DB.Where("username = ?", u.Username).First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}
	return &user, err
}

// ChangePassword 修改密码
func (UserService *UserService) ChangePassword(u *models.User, newPassword string) (userInfo *models.User, err error) {
	var user models.User
	if err = global.CASBIN_DB.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	user.Password = utils.BcryptHash(newPassword)
	err = global.CASBIN_DB.Save(&user).Error
	return &user, nil
}

// GetUserInfoList 获取用户列表信息
func (UserService *UserService) GetUserInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.CASBIN_DB.Model(&models.User{})
	var userList []models.User
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&userList).Error
	return userList, total, err
}

// DeleteUser 通过 ID 删除用户
func (UserService *UserService) DeleteUser(id int) (err error) {
	var user models.User
	err = global.CASBIN_DB.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	return err
}

// FindUserById 通过id 获取用户信息以及相关权限信息
func (UserService *UserService) FindUserById(id int) (user *models.User, err error) {
	var u models.User
	err = global.CASBIN_DB.Where("`id` = ?", id).First(&u).Error
	return &u, err
}
