package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wang900115/Perry/internal/adapter/controller"
)

type ToDo struct {
	todo controller.ToDo
}

func NewToDoRouter(todo *controller.ToDo) IRoute {
	return &ToDo{todo: *todo}
}

func (t *ToDo) SetUp(router *gin.RouterGroup) {
	todoGroup := router.Group("v1/todo")
	{
		todoGroup.POST("/dispatch", t.todo.Create)
		todoGroup.DELETE("/delete", t.todo.Delete)
		todoGroup.PUT("/update", t.todo.Update)
		todoGroup.GET("/query", t.todo.Query)
	}
}
