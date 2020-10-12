package route

import (
	"github.com/gin-gonic/gin"
	"github.com/sarailQAQ/destroy-demo/middleware"
	"github.com/sarailQAQ/destroy-demo/service"
)

//relative path: /comment
func commentRoute(r *gin.RouterGroup) {
	r.POST("/addComment", middleware.LoginStatus, service.PostComment)
	r.GET("/getCommentList", service.GetCommentList)
}
