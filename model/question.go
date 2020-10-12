package model

import (
	"github.com/jinzhu/gorm"
)

type QuestionTable struct {
	gorm.Model
	Tittle string
	Description string
	Uid uint
	Place string
	AnswerNum int
}

type QuestionImageTable struct {
	gorm.Model
	QuestionId uint
	ImageId uint
}

type QuestionJSON struct {
	QuestionId uint `json:"question_id"`
	Tittle string `json:"tittle"`
	Description string `json:"description"`
	Place string `json:"place"`
	Nickname string `json:"nickname"`
	PhotoAvatar string `json:"photo_avatar"`
	PhotoUrl []string `json:"photo_url"`
	AnswerNum int `json:"answer_num"`
	IsLike int `json:"is_like"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
