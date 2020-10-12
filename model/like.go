package model

import "github.com/jinzhu/gorm"

const (
	QuestionLike = 1
	AnswerLike = 2
)

type LikeTable struct {
	gorm.Model
	Type int
	Uid uint
	TableId uint
	Status int
}

type LikeJSON struct {
	 Status int `json:"status"`
}
