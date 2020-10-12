package dao

import "github.com/sarailQAQ/destroy-demo/model"

func GetAnswerLists(questionId uint) (answerTables []model.AnswerTable, err error) {
	err = DB.Model(&model.AnswerTable{}).Where("question_id=?", questionId).Find(&answerTables).Error
	return answerTables, err
}

func GetAnswerImage(answerId uint) (images []uint, err error) {
	var answerImages []model.AnswerImageTable
	err = DB.Model(&model.AnswerImageTable{}).Where("answer_id=?", answerId).Find(&answerImages).Error
	if err != nil {
		return nil, err
	}
	for _, image := range answerImages {
		images = append(images, image.ImageId)
	}
	return images, nil
}

func PostAnswer(table model.AnswerTable) (uint, error) {
	var question model.QuestionTable
	DB.Model(&model.QuestionTable{}).Where("id=?", table.QuestionId).First(&question)
	DB.Model(&model.QuestionTable{}).Where("id=?", table.QuestionId).Update("answer_num", question.AnswerNum + 1)
	err := DB.Create(&table).Error
	return table.ID, err
}

func PostAnswerImage(table model.AnswerImageTable) error {

	return DB.Create(&table).Error
}
