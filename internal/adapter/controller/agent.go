package controller

import (
	"github.com/gin-gonic/gin"
	responser "github.com/wang900115/Perry/internal/adapter/response"
	"github.com/wang900115/Perry/internal/adapter/validator"
	"github.com/wang900115/Perry/internal/application/usecase"
)

type Agent struct {
	agent    usecase.Agent
	response responser.Response
}

func NewAgentController(agent *usecase.Agent, response responser.Response) *Agent {
	return &Agent{agent: *agent, response: response}
}

// 新增代理(需用middleware帶參數)
func (a *Agent) Add(c *gin.Context) {
	var req validator.AgentAddRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		a.response.ClientFail400(c, err)
		return
	}
	// 為 ctx value
	userID := c.GetUint("user_id")
	agent, err := a.agent.Add(c, userID, req)
	if err != nil {
		a.response.ServerFail500(c, err)
		return
	}
	a.response.Success201(c, map[string]interface{}{
		"agent": agent,
	})
}

// 移除代理
func (a *Agent) Remove(c *gin.Context) {
	var req validator.AgentRemoveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		a.response.ClientFail400(c, err)
		return
	}
	err := a.agent.Remove(c, req)
	if err != nil {
		a.response.ServerFail500(c, err)
		return
	}
	a.response.Success204(c)
}

// 移除該使用者底下的代理(需用middleware帶參數)
func (a *Agent) RemoveAll(c *gin.Context) {
	userId := c.GetUint("user_id")
	err := a.agent.RemoveAll(c, userId)
	if err != nil {
		a.response.ServerFail500(c, err)
		return
	}
	a.response.Success204(c)
}

// 取得該使用者底下的代理(需用middleware帶參數)
func (a *Agent) Read(c *gin.Context) {
	userId := c.GetUint("user_id")
	agents, err := a.agent.Read(c, userId)
	if err != nil {
		a.response.ServerFail500(c, err)
		return
	}
	a.response.Success200(c, map[string]interface{}{
		"agents": agents,
	})
}
