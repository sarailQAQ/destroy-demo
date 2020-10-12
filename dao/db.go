package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/sarailQAQ/destroy-demo/model"
	"log"
)

var DB  *gorm.DB

func init() {
	sql, err := gorm.Open("mysql", "root:weweixiao228@tcp(127.0.0.1:3306)/destroy?charset=utf8&parseTime=true")
	if err != nil {
		log.Println("dao-db: ", err)
		return
	}
	DB = sql

	if !DB.HasTable(&model.UserTable{}) {
		DB.CreateTable(&model.UserTable{})
	} else {
		DB.AutoMigrate(&model.UserTable{})
	}

	if !DB.HasTable(&model.ImageTable{}) {
		DB.CreateTable(&model.ImageTable{})
	} else {
		DB.AutoMigrate(&model.ImageTable{})
	}

	if !DB.HasTable(&model.TipTable{}) {
		DB.CreateTable(&model.TipTable{})
	} else {
		DB.AutoMigrate(&model.TipTable{})
	}

	if !DB.HasTable(&model.AnswerTable{}) {
		DB.CreateTable(&model.AnswerTable{})
	} else {
		DB.AutoMigrate(&model.AnswerTable{})
	}

	if !DB.HasTable(&model.AnswerImageTable{}) {
		DB.CreateTable(&model.AnswerImageTable{})
	} else {
		DB.AutoMigrate(&model.AnswerImageTable{})
	}

	if !DB.HasTable(&model.QuestionTable{}) {
		DB.CreateTable(&model.QuestionTable{})
	} else {
		DB.AutoMigrate(&model.QuestionTable{})
	}

	if !DB.HasTable(&model.QuestionImageTable{}) {
		DB.CreateTable(&model.QuestionImageTable{})
	} else {
		DB.AutoMigrate(&model.QuestionImageTable{})
	}

	if !DB.HasTable(&model.CommentTable{}) {
		DB.CreateTable(&model.CommentTable{})
	} else {
		DB.AutoMigrate(&model.CommentTable{})
	}

	if !DB.HasTable(&model.CollectionTable{}) {
		DB.CreateTable(&model.CollectionTable{})
	} else {
		DB.AutoMigrate(&model.CollectionTable{})
	}

	if !DB.HasTable(&model.LikeTable{}) {
		DB.CreateTable(&model.LikeTable{})
	} else {
		DB.AutoMigrate(&model.LikeTable{})
	}
}
