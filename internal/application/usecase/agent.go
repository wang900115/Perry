package usecase

import (
	"context"

	"github.com/wang900115/Perry/internal/adapter/validator"
	"github.com/wang900115/Perry/internal/domain/entity"
	gorminterface "github.com/wang900115/Perry/internal/domain/interface/gorm"
)

type Agent struct {
	// db
	agentRepo gorminterface.Agent
	// !todo redis
}

func NewAgentUsecase(agentRepo *gorminterface.Agent) *Agent {
	return &Agent{agentRepo: *agentRepo}
}

func (a *Agent) Add(ctx context.Context, user_id uint, input validator.AgentAddRequest) (*entity.Agent, error) {
	return a.agentRepo.Add(ctx, user_id, input)
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
