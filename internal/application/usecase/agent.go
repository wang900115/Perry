package usecase

import (
	"context"

	"github.com/wang900115/Perry/internal/adapter/validator"
	"github.com/wang900115/Perry/internal/domain/entity"
	gorminterface "github.com/wang900115/Perry/internal/domain/interface/gorm"
	redisinterface "github.com/wang900115/Perry/internal/domain/interface/redis"
)

type Agent struct {
	// db
	agentRepo gorminterface.Agent
	// redis
	agentCache redisinterface.Agent
}

func NewAgentUsecase(agentRepo *gorminterface.Agent) *Agent {
	return &Agent{agentRepo: *agentRepo}
}

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

func (a *Agent) Remove(ctx context.Context, input validator.AgentRemoveRequest) error {
	return a.agentRepo.Remove(ctx, input)
}

func (a *Agent) RemoveAll(ctx context.Context, user_id uint) error {
	return a.agentRepo.RemoveAll(ctx, user_id)
}

func (a *Agent) Read(ctx context.Context, user_id uint) ([]*entity.Agent, error) {
	return a.agentRepo.Read(ctx, user_id)
}
