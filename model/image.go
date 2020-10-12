package model

import "github.com/jinzhu/gorm"

type ImageTable struct {
	gorm.Model
	Url string
}
