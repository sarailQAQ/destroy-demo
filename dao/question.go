package dao

import "github.com/sarailQAQ/destroy-demo/model"

func GetQuestionInfo(questionId uint) (model.QuestionTable, error) {
	var questionTable model.QuestionTable
	err := DB.Model(&questionTable).Where("id = ?", questionId).First(&questionTable).Error
	return questionTable, err
}

func GetQuestionImage(questionId uint) ([]string, error) {
	var questionImagesId []model.QuestionImageTable
	DB.Model(&model.QuestionImageTable{}).Select("image_id").Where("question_id=?", questionId).Find(&questionImagesId)

	var res = make([]string, 0)
	for _, table := range questionImagesId {
		var imageTable model.ImageTable
		DB.Model(&model.ImageTable{}).Select("url").Where("id=?",table.ImageId).First(&imageTable)
		res = append(res, imageTable.Url)
	}

	return res, nil
}

func PostQuestion(table model.QuestionTable) (uint, error) {
	err := DB.Create(&table).Error
	return table.ID, err
}

func PostQuestionImage(table model.QuestionImageTable) error {
	return DB.Create(&table).Error
}

func SearchQuestion(key string) (res []uint) {
	var tables []model.QuestionTable
	DB.Model(&model.QuestionTable{}).Select("id").Where("tittle like ?", key + "%").Find(&tables)
	for _, table := range tables{
		res = append(res, table.ID)
	}
	return res
}

func GetCollection(uid uint) (question_ids []uint) {
	var tables []model.CollectionTable
	DB.Model(&model.CollectionTable{}).Where("uid=? and status=1", uid).Find(&tables)
	for _, table := range tables{
		question_ids = append(question_ids, table.QuestionId)
	}
	return question_ids
}

func GetQuestionListLatest(page int) (res []uint, err error) {
	var list []model.QuestionTable
	err = DB.Model(&model.QuestionTable{}).Order("created_at").Limit( 10).Offset((page - 1) * 10).Find(&list).Error
	for _, table := range list {
		res = append(res, table.ID)
	}
	return res, err
}
