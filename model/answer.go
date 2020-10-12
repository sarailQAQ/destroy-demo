package model

import "github.com/jinzhu/gorm"

type  AnswerTable struct {
	gorm.Model
	Description string
	QuestionId uint
	Uid uint
}

type AnswerImageTable struct {
	gorm.Model
	AnswerId uint
	ImageId uint
}

type AnswerJSON struct {
	AnswerId uint `json:"answer_id"`
	Description string `json:"description"`
	Nickname string `json:"nickname"`
	PhotoAvatar string `json:"photo_avatar"`
	PhotoUrl []string `json:"photo_url"`
	IsLike int `json:"is_like"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
