package model

import "github.com/jinzhu/gorm"

type CommentTable struct {
	gorm.Model
	Description string
	AnswerId uint
	Uid uint
}

type CommentJSON struct {
	CommentId uint `json:"comment_id"`
	Description string `json:"description"`
	Nickname string `json:"nickname"`
	PhotoAvatar string `json:"photo_avatar"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}