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

func PostComment(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		resp.Error(c, resp.ClientError, errors.New("token error"))
		return
	}
	u := user.(middleware.UserClaim)

	var table model.CommentTable
	table.Description = c.PostForm("description")
	answerIdString := c.PostForm("answer_id")
	t, _ := strconv.Atoi(answerIdString)
	table.AnswerId = uint(t)
	table.Uid = u.Id

	_, err := dao.PostComment(table)
	if err != nil {
		resp.Error(c, resp.MiddlewareError, err)
		return
	}
	resp.Ok(c)
}

func GetCommentList(c *gin.Context) {
	answerIdString := c.PostForm("answer_id")
	t, err := strconv.Atoi(answerIdString)
	if err != nil {
		resp.Error(c, resp.ClientError, err)
		return
	}
	answerId := uint(t)

	tables, _ := dao.GetCommentList(answerId)
	comments := make([]model.CommentJSON, 0)
	for _, table := range tables {
		user, e := dao.GetUserInfoById(table.Uid)
		if e != nil {
			continue
		}
		avatar, _ := dao.GetImageUrl(user.AvatarId)
		comment := model.CommentJSON{
			CommentId:   table.ID,
			Description: table.Description,
			Nickname:    user.Nickname,
			PhotoAvatar: avatar,
			CreatedAt:   table.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   table.UpdatedAt.Format("2006-01-02 15:04:05"),
//			DeletedAt:   table.DeletedAt.Format("2006-01-02 15:04:05"),
		}
		comments = append(comments, comment)
	}
	resp.OkWithData(c, resp.Success, comments)
}
