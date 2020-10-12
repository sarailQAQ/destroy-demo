package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sarailQAQ/destroy-demo/dao"
	"github.com/sarailQAQ/destroy-demo/middleware"
	"github.com/sarailQAQ/destroy-demo/model"
	"github.com/sarailQAQ/destroy-demo/utils/resp"
	"log"
	"strconv"
)

func PostLike(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		resp.Error(c, resp.ClientError, errors.New("token error"))
		return
	}
	u := user.(middleware.UserClaim)
	typString := c.PostForm("type")
	var typ int
	if typString == "question" {
		typ = model.QuestionLike
	} else if typString == "answer" {
		typ = model.AnswerLike
	} else {
		resp.Error(c, resp.ClientError, errors.New("like type error"))
		return
	}
	tableIdString := c.PostForm("id")
	tableId, _ := strconv.Atoi(tableIdString)

	status, err := dao.PostLike(model.LikeTable{
		Model:   gorm.Model{},
		Type:    typ,
		Uid:     u.Id,
		TableId: uint(tableId),
	})

	if err != nil {
		resp.Error(c, resp.MiddlewareError, err)
		log.Println(err)
		return
	}
	resp.OkWithData(c, resp.Success, model.LikeJSON{Status: status})
}
