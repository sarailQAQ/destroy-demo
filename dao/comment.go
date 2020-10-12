package dao

import "github.com/sarailQAQ/destroy-demo/model"

func PostComment(table model.CommentTable) (uint, error) {
	err := DB.Create(&table).Error
	return table.ID, err
}

func GetCommentList(answerId uint) (tables []model.CommentTable, err error) {
	err = DB.Model(&model.CommentTable{}).Where("answer_id=?", answerId).Find(&tables).Error
	return tables, err
}
