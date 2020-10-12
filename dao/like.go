package dao

import (
	"github.com/sarailQAQ/destroy-demo/model"
)

func PostLike(like model.LikeTable) (int, error) {
	// 是否被重复点赞
	var table model.LikeTable
	ok := DB.Model(&model.LikeTable{}).Where("uid=? and type=? and table_id=?", like.Uid, like.Type, like.TableId).First(&table).RecordNotFound()
	if !ok {
		db := DB.Model(&model.LikeTable{}).Where("uid=? and type=? and table_id=?", like.Uid, like.Type, like.TableId).Update("status", 1 - table.Status)
		return 1 - table.Status, db.Error

	}
	like.Status = 1
	return 1, DB.Create(&like).Error
}

func IsLike(like model.LikeTable) bool {
	// 是否被点赞
	var table model.LikeTable
	DB.Model(&model.LikeTable{}).Where("type=? and uid=? and table_id=? and status=1", like.Type,like.Uid,like.TableId).First(&table)
	return table.ID > 0
}
