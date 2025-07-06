package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wang900115/Perry/internal/adapter/controller"
)

type User struct {
	user controller.User
}

func NewUserRouter(user controller.User) IRoute {
	return &User{user: user}
}

func (u *User) SetUp(router *gin.RouterGroup) {
	userGroup := router.Group("v1/user")
	{
		userGroup.POST("/regist", u.user.Regist)
		userGroup.PUT("/update", u.user.Update)
		userGroup.DELETE("/delete", u.user.Delete)

		userGroup.POST("/login", u.user.Login)
		userGroup.POST("/logout", u.user.Logout)

		userGroup.POST("/forgot", u.user.Forgot)

	}
}
