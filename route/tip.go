package route

import (
	"github.com/gin-gonic/gin"
	"github.com/sarailQAQ/destroy-demo/service"
)

func tipRoute(r *gin.RouterGroup) {
	r.POST("/getTips", service.GetTips)
}
