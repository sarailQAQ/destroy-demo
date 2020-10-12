package route

import (
	"github.com/gin-gonic/gin"
	"log"
)

func SetRoute() {
	r := gin.Default()

	r.Static("/source/images", "./source/images")

	pictureRoute(&r.RouterGroup)

	userRoute(&r.RouterGroup)
	tipRoute(&r.RouterGroup)

	questionRoute(r.Group("/question"))
	answerRoute(r.Group("/answer"))
	commentRoute(r.Group("/comment"))
	likeRoute(&r.RouterGroup)
	collectionRoute(&r.RouterGroup)
	if err := r.Run(":8081"); err != nil {
		log.Println(err)
	}
}