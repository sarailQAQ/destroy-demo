package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sarailQAQ/destroy-demo/dao"
	"github.com/sarailQAQ/destroy-demo/model"
	"github.com/sarailQAQ/destroy-demo/utils/resp"
)

func GetTips(c *gin.Context)  {
	tittle := c.PostForm("tittle")
	if len(tittle) == 0 {
		resp.Error(c, resp.ClientError, errors.New("tittle not found"))
		return
	}

	tipTables, err := dao.GetTips(tittle)
	if err != nil {
		resp.Error(c, resp.ClientError, errors.New("tittle not found"))
	}
	tipJSONs := make([]model.TipJSON, 0)
	for _, tip := range tipTables {
		tipJSONs = append(tipJSONs, model.TipJSON{
			Id: tip.ID,
			Description: tip.Description,
		})
	}
	resp.OkWithData(c, resp.Success, tipJSONs)
}
