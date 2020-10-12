package dao

import "github.com/sarailQAQ/destroy-demo/model"

func GetTips(title string) ([]model.TipTable, error) {
	var Tips []model.TipTable
	err := DB.Model(&model.TipTable{}).Where("tittle = ?", title).Find(&Tips).Error
	if err != nil  {
		return nil, err
	}
	return Tips, nil
}
