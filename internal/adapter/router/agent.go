package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wang900115/Perry/internal/adapter/controller"
)

type Agent struct {
	agent controller.Agent
}

func NewAgentRouter(agent *controller.Agent) IRoute {
	return &Agent{agent: *agent}
}

func (a *Agent) SetUp(router *gin.RouterGroup) {
	agentGroup := router.Group("v1/agent")
	{
		agentGroup.POST("/add", a.agent.Add)
		agentGroup.DELETE("/remove", a.agent.Remove)
		agentGroup.DELETE("/remove/all", a.agent.RemoveAll)
		agentGroup.GET("/read", a.agent.Read)
	}
}
