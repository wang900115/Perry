package controller

import (
	"github.com/gin-gonic/gin"
	responser "github.com/wang900115/Perry/internal/adapter/response"
	"github.com/wang900115/Perry/internal/adapter/validator"
	"github.com/wang900115/Perry/internal/application/usecase"
)

type User struct {
	user     usecase.User
	response responser.Response
}

func NewUserController(user *usecase.User, response responser.Response) *User {
	return &User{user: *user, response: response}
}

// 註冊帳號密碼(不需要middleware帶參數)
func (u *User) Regist(c *gin.Context) {
	var req validator.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		u.response.ClientFail400(c, err)
		return
	}
	if err := u.user.Register(c, req); err != nil {
		u.response.ServerFail500(c, err)
		return
	}
	u.response.Success204(c)
}

// 單點登入(不需要middleware帶參數)
func (u *User) Login(c *gin.Context) {
	var req validator.LoginRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		u.response.ClientFail400(c, err)
		return
	}
	ip := c.ClientIP()
	userAgent := c.Request.UserAgent()
	user, token, err := u.user.Login(c, req.Username, req.Password, ip, userAgent)
	if err != nil {
		u.response.ServerFail500(c, err)
		return
	}
	u.response.Success200(c, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}

// 單點登出(需要middleware帶參數(user_id, session_id))
func (u *User) Logout(c *gin.Context) {
	// 為 ctx value
	userID := c.GetUint("user_id")
	sessionID := c.GetInt64("session_id")

	if err := u.user.Logout(c, userID, sessionID); err != nil {
		u.response.ServerFail500(c, err)
		return
	}
	u.response.Success204(c)
}

// 刪除帳號(需要middleware帶參數(user_id))
func (u *User) Delete(c *gin.Context) {
	// 為 ctx value
	userID := c.GetUint("user_id")
	if err := u.user.Delete(c, userID); err != nil {
		u.response.ServerFail500(c, err)
		return
	}
	u.response.Success204(c)
}

// 更新密碼(需要middleware帶參數(user_id))
func (u *User) UpdatePassword(c *gin.Context) {
	var req validator.UpdatePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		u.response.ClientFail400(c, err)
		return
	}
	// 為 ctx value
	userID := c.GetUint("user_id")
	if err := u.user.UpdatePassword(c, userID, req.Password); err != nil {
		u.response.ServerFail500(c, err)
		return
	}
	u.response.Success204(c)
}

// 更新設定(需要middleware帶參數(user_id))
func (u *User) UpdateSettings(c *gin.Context) {
	var req validator.UpdateSettingsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		u.response.ClientFail400(c, err)
		return
	}
	// 為 ctx value
	userID := c.GetUint("user_id")
	user, err := u.user.UpdateSettins(c, userID, req)
	if err != nil {
		u.response.ServerFail500(c, err)
		return
	}
	u.response.Success200(c, map[string]interface{}{
		"user": user,
	})
}

// !todo 找回帳號
func (u *User) Forgot(c *gin.Context) {
}
