package route

import (
	"github.com/gin-gonic/gin"
	"github.com/sarailQAQ/destroy-demo/middleware"
	"github.com/sarailQAQ/destroy-demo/service"
)

// relativePath: /question
func questionRoute(r *gin.RouterGroup)  {
	r.GET("/getQuestionInfo", service.GetQuestionInfo)
	r.POST("/search", service.SearchQuestion)
	r.GET("/getCollectedQuestion", middleware.LoginStatus, service.GetCollectQuestion)
	r.GET("/getQuestionListLatest", middleware.LoginStatus, service.GetQuestionListLatest)

	r.POST("/addQuestion", middleware.LoginStatus, middleware.UploadSeveralImage, service.PostQuestion)
}
