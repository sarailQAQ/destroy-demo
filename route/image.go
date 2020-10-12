package route

import (
	"github.com/gin-gonic/gin"
	"github.com/sarailQAQ/destroy-demo/service"
)

func pictureRoute(r *gin.RouterGroup) {
	r.POST("/upload", service.UploadImage)
}
