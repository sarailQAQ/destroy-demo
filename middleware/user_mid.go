package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sarailQAQ/destroy-demo/utils/resp"
	"github.com/sarailQAQ/destroy-demo/utils/token"
)

type UserClaim struct {
	Id uint
	Username,UserId string
}

func LoginStatus(c *gin.Context) {
	toke := c.Request.Header.Get("Authorization")
	if len(toke) < 7 {
		resp.Error(c, resp.ClientError, errors.New("Illegal jwt"))
		c.Abort()
		return
	}
	auth := toke[7:]
	uid,user,userId, err := token.Check(auth)
	if err != nil {
		resp.Error(c, resp.ClientError, err)
		c.Abort()
		return
	}

	c.Set("user",UserClaim{
		Id:       uid,
		Username: user,
		UserId: userId,
	})

	c.Next()
	return
}

func UserStatusBear(c *gin.Context) {
	toke := c.Request.Header.Get("Authorization")
	if len(toke) < 7 {
		resp.Error(c, resp.ClientError, errors.New("Illegal jwt"))
		c.Next()
		return
	}
	auth := toke[7:]
	uid,user,userId, err := token.Check(auth)
	if err != nil {
		resp.Error(c, resp.ClientError, err)
		c.Next()
		return
	}

	c.Set("user",UserClaim{
		Id:       uid,
		Username: user,
		UserId: userId,
	})

	c.Next()
}