package model

import "github.com/jinzhu/gorm"

type TipTable struct {
	gorm.Model
	Tittle string
	Description string
}

type TipJSON struct {
	Id uint `json:"id"`
	Description string `json:"description"`
}


