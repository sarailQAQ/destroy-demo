package model

import "github.com/jinzhu/gorm"

type CollectionTable struct {
	gorm.Model
	Uid uint
	QuestionId uint
	Status int
}

type CollectionJSON struct {
	Status int `json:"status"`
}
