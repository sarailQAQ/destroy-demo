package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sarailQAQ/destroy-demo/dao"
	"github.com/sarailQAQ/destroy-demo/middleware"
	"github.com/sarailQAQ/destroy-demo/model"
	"github.com/sarailQAQ/destroy-demo/utils/resp"
	"strconv"
)

func getQuestion(id uint, uid uint) (model.QuestionJSON, error) {
	quesTable, err := dao.GetQuestionInfo(uint(id))
	if err != nil {
		return model.QuestionJSON{}, err
	}
	urls, _ := dao.GetQuestionImage(uint(id))
	user, err := dao.GetUserInfoById(quesTable.Uid)
	if err != nil {
		return model.QuestionJSON{}, err
	}
	avatar, _ := dao.GetImageUrl(user.AvatarId)
	isSelf := 0
	if dao.IsLike(model.LikeTable{
		Type:    model.QuestionLike,
		Uid:     uid,
		TableId: id,
	}) {
		isSelf = 1
	}

	return model.QuestionJSON{
		QuestionId:  quesTable.ID,
		Tittle:      quesTable.Tittle,
		Description: quesTable.Description,
		Place:       quesTable.Place,
		Nickname:    user.Nickname,
		PhotoAvatar: avatar,
		PhotoUrl:    urls,
		IsLike: 	 isSelf,
		AnswerNum:   quesTable.AnswerNum,
		CreatedAt:   quesTable.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   quesTable.UpdatedAt.Format("2006-01-02 15:04:05"),
//		DeletedAt:   quesTable.DeletedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func GetQuestionInfo(c *gin.Context)  {
	questionId := c.PostForm("question_id")
	qIdInt, _ := strconv.Atoi(questionId)

	user, ok := c.Get("user")
	var uid uint
	if !ok {	uid = 0	} else {
		user := user.(middleware.UserClaim)
		uid  = user.Id
	}

	res, err := getQuestion(uint(qIdInt), uid)
	if err != nil {
		resp.Error(c, resp.ClientError, err)
		return
	}
	resp.OkWithData(c, resp.Success, res)
}

func SearchQuestion(c *gin.Context) {
	key := c.PostForm("tittle")
	qIds := dao.SearchQuestion(key)
	user, ok := c.Get("user")
	var uid uint
	if !ok {	uid = 0	} else {
		user := user.(middleware.UserClaim)
		uid  = user.Id
	}

	questions := make([]model.QuestionJSON, 0)
	for _, qid := range qIds {
		question, _ := getQuestion(qid, uid)
		questions = append(questions, question)
	}

	resp.OkWithData(c, resp.Success, questions)
}

func GetCollectQuestion(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		resp.Error(c, resp.ClientError, errors.New("token error"))
		return
	}
	u := user.(middleware.UserClaim)

	questionIds := dao.GetCollection(u.Id)
	questions := make([]model.QuestionJSON, 0)
	for _, qid := range questionIds {
		question, _ := getQuestion(qid, u.Id)
		questions = append(questions, question)
	}
	resp.OkWithData(c, resp.Success, questions)
}

func PostQuestion(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		resp.Error(c, resp.ClientError, errors.New("token error"))
		return
	}
	u := user.(middleware.UserClaim)

	var qTable model.QuestionTable
	qTable.Tittle = c.PostForm("tittle")
	qTable.Place = c.PostForm("place")
	qTable.Description = c.PostForm("description")
	qTable.Uid = u.Id
	var err error
	qTable.ID, err = dao.PostQuestion(qTable)
	if err != nil {
		resp.Error(c, resp.ClientError, err)
		return
	}

	count, _ := c.Get("count")

	for i := 1; i <= count.(int); i++ {
		filename, ok := c.Get("photo" + strconv.Itoa(i))
		if ok {
			imagID, err := dao.UploadImage(filename.(string))
			if err == nil {
				_ = dao.PostQuestionImage(model.QuestionImageTable{
					QuestionId: qTable.ID,
					ImageId:    imagID,
				})
			}
		}
	}
	resp.Ok(c)
}

func GetQuestionListLatest(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		resp.Error(c, resp.ClientError, errors.New("token error"))
		return
	}
	u := user.(middleware.UserClaim)
	page_s := c.PostForm("page")
	page, err := strconv.Atoi(page_s)
	if err != nil {
		resp.Error(c, resp.ClientError, err)
		return
	}
	
	ids, _ := dao.GetQuestionListLatest(page)
	questions := make([]model.QuestionJSON, 0)
	for _, id := range ids {
		qJSON, _ := getQuestion(id, u.Id)
		questions = append(questions, qJSON)
	}

	resp.OkWithData(c, resp.Success, questions)
}