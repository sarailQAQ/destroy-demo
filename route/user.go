package route

import (
	"github.com/gin-gonic/gin"
	"github.com/sarailQAQ/destroy-demo/middleware"
	"github.com/sarailQAQ/destroy-demo/service"
)

func userRoute(r *gin.RouterGroup) {
	r.POST("/signin", service.SignIn)
	r.POST("/signup", service.SignUp)

	v := r.Group("/user")
	{
		v.Use(middleware.LoginStatus)
		v.POST("/getUserInfo", service.GetUserInfo)
		v.POST("/updateUserInfo", middleware.UploadSingelImage, service.UpdateUserInfo)
	}
}
