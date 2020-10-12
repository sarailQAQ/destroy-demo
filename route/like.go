package route

import (
	"github.com/gin-gonic/gin"
	"github.com/sarailQAQ/destroy-demo/middleware"
	"github.com/sarailQAQ/destroy-demo/service"
)

// relative path: /
func likeRoute(r *gin.RouterGroup) {
	r.POST("/like", middleware.LoginStatus, service.PostLike)
}
