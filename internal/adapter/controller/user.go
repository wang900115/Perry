package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wang900115/Perry/internal/application/usecase"
)

type User struct {
	user usecase.User
}

func NewUserController(user *usecase.User) *User {
	return &User{user: *user}
}

func (u *User) Regist(c *gin.Context) {

}

func (u *User) Update(c *gin.Context) {

}

func (u *User) Login(c *gin.Context) {

}

func (u *User) Logout(c *gin.Context) {

}

func (u *User) Forgot(c *gin.Context) {

}

func (u *User) Delete(c *gin.Context) {

}
