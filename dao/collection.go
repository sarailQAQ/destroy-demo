package dao

import "github.com/sarailQAQ/destroy-demo/model"

func PostCollection(collection model.CollectionTable) (int, error) {
	// 是否被重复收藏
	var table model.CollectionTable
	ok := DB.Model(&model.CollectionTable{}).Where("uid=? and question_id=?", collection.Uid,  collection.QuestionId).First(&table).RecordNotFound()
	if !ok {
		db := DB.Model(&model.CollectionTable{}).Where("uid=? and question_id=?", collection.Uid,  collection.QuestionId)
		return 1 - table.Status, db.Update("status", 1 - table.Status).Error
	}
	collection.Status = 1
	return 1, DB.Create(&collection).Error
}
