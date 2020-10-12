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

func GetAnswerList(c *gin.Context) {
	questionIdString := c.Request.FormValue("question_id")
	questionId, err := strconv.Atoi(questionIdString)
	if err != nil {
		resp.Error(c, resp.ClientError, err)
		return
	}

	tables, _ := dao.GetAnswerLists(uint(questionId))

	uid := uint(0)
	user, ok := c.Get("user")
	if ok {
		u := user.(middleware.UserClaim)
		uid = u.Id
	}

	data := make([]model.AnswerJSON, 0)
	for _, table := range tables {
		userTable, err := dao.GetUserInfoById(uid)
		if err != nil {
			continue
		}
		avatar, _ := dao.GetImageUrl(userTable.AvatarId)
		imageIds, _ := dao.GetAnswerImage(table.ID)
		urls := dao.GetSeveralImages(imageIds)
		isLike := 0
		if dao.IsLike(model.LikeTable{
			Type:    model.AnswerLike,
			Uid:     uid,
			TableId: table.ID,
		}) {	isLike = 1	}
		answer := model.AnswerJSON{
			AnswerId:    table.ID,
			Description: table.Description,
			Nickname:    userTable.Nickname,
			PhotoAvatar: avatar,
			PhotoUrl:    urls,
			IsLike: 	 isLike,
			CreatedAt:   table.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   table.UpdatedAt.Format("2006-01-02 15:04:05"),
//			DeletedAt:   table.DeletedAt.Format("2006-01-02 15:04:05"),
		}
		data = append(data, answer)
	}
	resp.OkWithData(c, resp.Success, data)
}

func PostAnswer(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		resp.Error(c, resp.ClientError, errors.New("token error"))
		return
	}
	u := user.(middleware.UserClaim)

	var aTable model.AnswerTable
//	aTable.Tittle = c.PostForm("tittle")
//	aTable.Place = c.PostForm("place")
	aTable.Description = c.PostForm("description")
	questionId, err := strconv.Atoi(c.PostForm("question_id"))
	if err != nil {
		resp.Error(c, resp.ClientError, err)
		return
	}
	aTable.QuestionId = uint(questionId)
	aTable.Uid = u.Id
	aTable.ID, err = dao.PostAnswer(aTable)
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
				_ = dao.PostAnswerImage(model.AnswerImageTable{
					AnswerId: aTable.ID,
					ImageId:    imagID,
				})
			}
		}
	}
	resp.Ok(c)
}