package controller

import (
	"github.com/gin-gonic/gin"
	responser "github.com/wang900115/Perry/internal/adapter/response"
	"github.com/wang900115/Perry/internal/adapter/validator"
	"github.com/wang900115/Perry/internal/application/usecase"
)

type ToDo struct {
	todo     usecase.ToDo
	response responser.Response
}

func NewToDoController(todo *usecase.ToDo, response responser.Response) *ToDo {
	return &ToDo{todo: *todo, response: response}
}

// 新增待辦事項(需要middleware帶參數(user_id))
func (t *ToDo) Create(c *gin.Context) {
	var req validator.ToDoCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		t.response.ClientFail400(c, err)
		return
	}
	// 為 ctx value
	userID := c.GetUint("user_id")
	todo, err := t.todo.Create(c, userID, req)
	if err != nil {
		t.response.ServerFail500(c, err)
		return
	}
	t.response.Success201(c, map[string]interface{}{
		"todo": todo,
	})
}

// 更新待辦事項(需要middleware帶參數(user_id))
func (t *ToDo) Update(c *gin.Context) {
	var req validator.ToDoUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		t.response.ClientFail400(c, err)
		return
	}
	// 為 ctx value
	userID := c.GetUint("user_id")
	todo, err := t.todo.Update(c, userID, req)
	if err != nil {
		t.response.ServerFail500(c, err)
		return
	}
	t.response.Success200(c, map[string]interface{}{
		"todo": todo,
	})
}

// 刪除待辦事項(需要middleware帶參數(user_id))
func (t *ToDo) Delete(c *gin.Context) {
	var req validator.ToDoDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		t.response.ClientFail400(c, err)
		return
	}
	// 為 ctx value
	userID := c.GetUint("user_id")
	err := t.todo.Delete(c, userID, req)
	if err != nil {
		t.response.ServerFail500(c, err)
		return
	}
	t.response.Success204(c)
}

// 取得使用者待辦事項(需要middleware帶參數(user_id))
func (t *ToDo) Query(c *gin.Context) {
	// 為 ctx value
	userID := c.GetUint("user_id")
	todos, err := t.todo.Query(c, userID)
	if err != nil {
		t.response.ServerFail500(c, err)
		return
	}
	t.response.Success200(c, map[string]interface{}{
		"todos": todos,
	})
}
