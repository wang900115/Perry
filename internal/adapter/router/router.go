package router

import "github.com/gin-gonic/gin"

type IRoute interface {
	SetUp(router *gin.RouterGroup)
}
