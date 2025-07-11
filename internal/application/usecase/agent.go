package usecase

import (
	"context"

	"github.com/wang900115/Perry/internal/adapter/validator"
	"github.com/wang900115/Perry/internal/domain/entity"
	gorminterface "github.com/wang900115/Perry/internal/domain/interface/gorm"
	redisinterface "github.com/wang900115/Perry/internal/domain/interface/redis"
)

type Agent struct {
	agentRepo  gorminterface.Agent
	agentCache redisinterface.Agent
}

func NewAgentUsecase(agentRepo *gorminterface.Agent, agentCache *redisinterface.Agent) *Agent {
	return &Agent{agentRepo: *agentRepo, agentCache: *agentCache}
}

// 新增代理(先在db新增再redis初始化)
func (a *Agent) Add(ctx context.Context, user_id uint, input validator.AgentAddRequest) (*entity.Agent, error) {
	agent, err := a.agentRepo.Add(ctx, user_id, input)
	if err != nil {
		return nil, err
	}
	err = a.agentCache.Initialize(ctx, agent.ID)
	if err != nil {
		return nil, err
	}
	return agent, nil
}

// 移除代理(先移除db再移除redis)
func (a *Agent) Remove(ctx context.Context, input validator.AgentRemoveRequest) error {
	err := a.agentRepo.Remove(ctx, input)
	if err != nil {
		return err
	}
	return a.agentCache.Delete(ctx, input.ID)
}

// 移除使用者的全部代理(先移除db再移除redis)
func (a *Agent) RemoveAll(ctx context.Context, user_id uint) error {
	err := a.agentRepo.RemoveAll(ctx, user_id)
	if err != nil {
		return err
	}
	return a.agentCache.DeleteAll(ctx, user_id)
}

// 取得使用者的全部代理(先去redis檢查如果沒有則去db)
func (a *Agent) Read(ctx context.Context, user_id uint) ([]*entity.Agent, error) {
	agentModels, err := a.agentCache.Get(ctx, user_id)
	if err != nil {
		return nil, err
	}
	if len(agentModels) > 0 {
		agents := make([]*entity.Agent, 0, len(agentModels))
		for _, agentModel := range agentModels {
			agents = append(agents, agentModel.ToDomain())
		}
		return agents, nil
	}
	return a.agentRepo.Read(ctx, user_id)
}
