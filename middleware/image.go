package middleware

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
)


const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func UploadSingelImage(c *gin.Context) {
	file, err := c.FormFile("photo1")
	if err != nil {
		c.Next()
		return
	}
	filename := randStringBytes(32) + ".png"
	err = c.SaveUploadedFile(file, "./source/images/" + filename)
	if err != nil {
		c.Next()
		return
	}
	c.Set("photo1", "/source/images/" + filename)
}

func UploadSeveralImage(c *gin.Context) {
	var i int
	for i = 0; i < 10; i++ {
		key := "photo" + strconv.Itoa(i + 1)
		file, err := c.FormFile(key)
		if file == nil {
			break
		}
		if err != nil {
			c.Next()
			return
		}
		filename := randStringBytes(32) + ".png"
		err = c.SaveUploadedFile(file, "./source/images/" + filename)
		if err != nil {
			c.Next()
			return
		}
		c.Set(key, "/source/images/" + filename)
	}
	c.Set("count", i)
}
