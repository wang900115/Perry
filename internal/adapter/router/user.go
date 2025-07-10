package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wang900115/Perry/internal/adapter/controller"
	"github.com/wang900115/Perry/internal/adapter/middleware/jwt"
)

type User struct {
	user controller.User
	jwt  jwt.JWT
}

func NewUserRouter(user *controller.User, jwt *jwt.JWT) IRoute {
	return &User{user: *user, jwt: *jwt}
}

func (u *User) SetUp(router *gin.RouterGroup) {
	userGroup := router.Group("v1/user")
	{
		userGroup.POST("/regist", u.user.Regist)
		userGroup.PUT("/update/settings", u.jwt.Middleware, u.user.UpdateSettings)
		userGroup.PATCH("/update/password", u.jwt.Middleware, u.user.UpdatePassword)
		userGroup.DELETE("/delete", u.jwt.Middleware, u.user.Delete)
		userGroup.POST("/login", u.user.Login)
		userGroup.POST("/logout", u.jwt.Middleware, u.user.Logout)
		userGroup.POST("/forgot", u.user.Forgot)
	}
}
