package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type UserTable struct {
	gorm.Model
	UserId string `gorm:"primary_key"`
	Password string
	Nickname string
	AvatarId uint
	Description string
	Qq string
	Tel string
	UserStatus string
}

type UserJson struct {
	UserId uint `json:"user_id"`
	Nickname string `json:"nickname"`
	Description string `json:"description"`
	Qq string `json:"qq"`
	Tel string `json:"tel"`
	UserStatus string `json:"user_status"`
	Avatar string `json:"avatar"`
}