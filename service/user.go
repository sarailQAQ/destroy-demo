package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sarailQAQ/destroy-demo/dao"
	"github.com/sarailQAQ/destroy-demo/middleware"
	"github.com/sarailQAQ/destroy-demo/model"
	"github.com/sarailQAQ/destroy-demo/utils/resp"
	"github.com/sarailQAQ/destroy-demo/utils/token"
)

func GetUserInfo(c *gin.Context) {
	user, _ := c.Get("user")
	userClaim, _ := user.(middleware.UserClaim)
	userTable, err1 := dao.GetUserInfo(userClaim.UserId)
	if err1 != nil {
		resp.Error(c, resp.ClientError, err1)
		return
	}
	avatarUrl, _ := dao.GetImageUrl(userTable.AvatarId)

	resp.OkWithData(c, resp.Success, model.UserJson{
		UserId:      userTable.ID,
		Nickname:    userTable.Nickname,
		Description: userTable.Description,
		Qq:          userTable.Qq,
		Tel:         userTable.Tel,
		UserStatus:  userTable.UserStatus,
		Avatar: 	 avatarUrl,
	})
}

func UpdateUserInfo(c *gin.Context) {
	user, _ := c.Get("user")
	userClaim, _ := user.(middleware.UserClaim)
	userTable, err := dao.GetUserInfo(userClaim.UserId)
	if err != nil {
		resp.Error(c, resp.ClientError, err)
		return
	}

	var tmp string
	tmp = c.PostForm("nickname")
	if len(tmp) > 0 {	userTable.Nickname = tmp 	}
	tmp = c.PostForm("description")
	if len(tmp) > 0 {	userTable.Description = tmp 	}
	tmp = c.PostForm("qq")
	if len(tmp) > 0 {	userTable.Qq = tmp 	}
	tmp = c.PostForm("tel")
	if len(tmp) > 0 {	userTable.Tel = tmp 	}
	tmp = c.PostForm("user_status")
	if len(tmp) > 0 {	userTable.UserStatus = tmp 	}

	if url, ok := c.Get("photo1"); ok {
		url := url.(string)
		avatarId, err := dao.UploadImage(url)
		if err == nil {
			_ = dao.PostAvatar(userClaim.Id, avatarId)
		}
	}

	err = dao.UpdateUserInfo(userTable)
	if err != nil {
		resp.Error(c, resp.MiddlewareError, err)
		return
	}
	resp.OkWithData(c, resp.Success, gin.H{"message": "update success"})

}

func SignIn(c *gin.Context) {
	userId := c.PostForm("user_id")
	passwprd := c.PostForm("password")
	userTable := dao.SignIn(model.UserTable{
		UserId: userId,
		Password: passwprd,
	})
	if userTable.ID > 0 {
		toke := token.Creat(userTable.ID, userTable.Nickname, userId)
		resp.OkWithData(c, resp.Success, gin.H{
			"token":       toke,
			"userId":      userTable.ID,
			"nickname":     userTable.Nickname,
			"description": userTable.Description,
			"qq":          userTable.Qq,
			"tel":         userTable.Tel,
			"userStatus":  userTable.UserStatus,
		})
		return
	}
	resp.Error(c, resp.ClientError, errors.New("用户名密码错误"))
}

func SignUp(c *gin.Context) {
	userId := c.PostForm("user_id")
	passwprd := c.PostForm("password")
	status := c.DefaultPostForm("user_status", "100")
	err := dao.SignUp(&model.UserTable{
		UserId: userId,
		Password: passwprd,
		UserStatus: status,
		Nickname: userId,
	})
	if err != nil {
		resp.Error(c, resp.ClientError, err)
		return
	}
	resp.Ok(c)
}