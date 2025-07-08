package controller

import (
	"github.com/gin-gonic/gin"
	responser "github.com/wang900115/Perry/internal/adapter/response"
	"github.com/wang900115/Perry/internal/application/usecase"
)

type User struct {
	user     usecase.User
	response responser.Response
}

func NewUserController(user *usecase.User, response responser.Response) *User {
	return &User{user: *user, response: response}
}

func (u *User) Regist(c *gin.Context) {
}

func (u *User) Update(c *gin.Context) {
}

func (u *User) Login(c *gin.Context) {
}

func (u *User) Logout(c *gin.Context) {
}

func (u *User) Delete(c *gin.Context) {
}

// !todo 找回帳號
func (u *User) Forgot(c *gin.Context) {
}
