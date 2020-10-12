package route

import (
	"github.com/gin-gonic/gin"
	"github.com/sarailQAQ/destroy-demo/middleware"
	"github.com/sarailQAQ/destroy-demo/service"
)

// relative path: /
func collectionRoute(r *gin.RouterGroup) {
	r.POST("/collect", middleware.LoginStatus, service.PostCollection)
}
