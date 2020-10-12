package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sarailQAQ/destroy-demo/dao"
	"github.com/sarailQAQ/destroy-demo/middleware"
	"github.com/sarailQAQ/destroy-demo/model"
	"github.com/sarailQAQ/destroy-demo/utils/resp"
	"log"
	"strconv"
)

func PostCollection(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		resp.Error(c, resp.ClientError, errors.New("token error"))
		return
	}
	u := user.(middleware.UserClaim)
	tableIdString := c.PostForm("id")
	tableId, _ := strconv.Atoi(tableIdString)

	status, err := dao.PostCollection(model.CollectionTable{
		Uid:     u.Id,
		QuestionId: uint(tableId),
	})

	if err != nil {
		resp.Error(c, resp.MiddlewareError, err)
		log.Println(err)
		return
	}
	resp.OkWithData(c, resp.Success, model.CollectionJSON{Status: status})
}
