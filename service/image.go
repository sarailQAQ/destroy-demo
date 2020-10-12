package service

import (
	"github.com/gin-gonic/gin"
	"log"
)

func UploadImage(c *gin.Context) {
	file, _ := c.FormFile("file")
	name := c.PostForm("id")

	//filename := file.Filename
	filename := name + ".png"
	if err := c.SaveUploadedFile(file, "./source/images/" + filename); err != nil {
		c.JSON(200, gin.H{"code":2000, "message": "uppload failed"})
		log.Println(err)
		return
	}
	c.JSON(200, gin.H{"code": 1000,"message":"upload success"})
}
