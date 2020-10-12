package route

import (
	"github.com/gin-gonic/gin"
	"github.com/sarailQAQ/destroy-demo/middleware"
	"github.com/sarailQAQ/destroy-demo/service"
)

// relative path: /answer
func answerRoute(r *gin.RouterGroup)  {
	r.GET("/getAnswerList", middleware.UserStatusBear,service.GetAnswerList)
	r.POST("/addAnswer", middleware.LoginStatus, middleware.UploadSeveralImage, service.PostAnswer)
}
