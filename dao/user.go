package dao

import (
	"errors"
	"fmt"
	"github.com/sarailQAQ/destroy-demo/model"
)

func GetUserInfo(uid string) (model.UserTable, error) {
	userTable := model.UserTable{

	}
	err := DB.Model(&model.UserTable{}).Where("user_id=?",uid).First(&userTable).Error
	return userTable, err
}

func GetUserInfoById(id uint) (model.UserTable, error) {
	userTable := model.UserTable{}
	if err := DB.Model(&userTable).Where("id=?", id).First(&userTable).Error; err != nil {
		return model.UserTable{}, err
	}
	return userTable, nil
}

func UpdateUserInfo(userTable model.UserTable) error {
		return DB.Model(&userTable).Update(&userTable).Error
}

func SignUp(userTable *model.UserTable) error {
	var user model.UserTable
	ok := DB.Model(&model.UserTable{}).Where("user_id = ?", userTable.UserId).First(&user).RecordNotFound()
	if !ok {
		return errors.New("用户已存在")
	}

	return DB.Create(&userTable).Error
}

func SignIn(userTable model.UserTable) model.UserTable {
	var user model.UserTable
	DB.Model(&model.UserTable{}).Where("user_id=? and password=?",userTable.UserId, userTable.Password).First(&user)
	fmt.Println(user)
	return user
}

func PostAvatar(id uint, urlId uint) error {
	return DB.Model(&model.UserTable{}).Where("id = ?", id).Update("avatar_id", urlId).Error
}