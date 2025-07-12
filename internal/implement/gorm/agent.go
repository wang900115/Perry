package gormimplement

import (
	"context"

	gormmodel "github.com/wang900115/Perry/internal/adapter/gorm/model"
	"github.com/wang900115/Perry/internal/adapter/validator"
	"github.com/wang900115/Perry/internal/domain/entity"
	gorminterface "github.com/wang900115/Perry/internal/domain/interface/gorm"
	"gorm.io/gorm"
)

type Agent struct {
	gorm *gorm.DB
}

func NewAgentImplement(gorm *gorm.DB) gorminterface.Agent {
	return &Agent{gorm: gorm}
}

func (a *Agent) Add(ctx context.Context, user_id uint, input validator.AgentAddRequest) (*entity.Agent, error) {
	createdAgent := gormmodel.Agent{
		UserID:      user_id,
		Name:        input.Name,
		Age:         input.Age,
		Role:        input.Role,
		Language:    input.Language,
		Description: input.Description,
	}
	if err := a.gorm.WithContext(ctx).Create(&createdAgent).Error; err != nil {
		return nil, err
	}
	return createdAgent.Domain(), nil
}

func (a *Agent) Remove(ctx context.Context, input validator.AgentRemoveRequest) error {
	var agentModel gormmodel.Agent
	if err := a.gorm.WithContext(ctx).Delete(&agentModel, input.ID).Error; err != nil {
		return err
	}
	return nil
}

func (a *Agent) RemoveAll(ctx context.Context, user_id uint) error {
	if err := a.gorm.WithContext(ctx).Where("user_id = ?", user_id).Delete(&gormmodel.Agent{}).Error; err != nil {
		return err
	}
	return nil
}

func (a *Agent) Read(ctx context.Context, user_id uint) ([]*entity.Agent, error) {
	var agentModels []gormmodel.Agent
	if err := a.gorm.WithContext(ctx).Where("user_id = ?", user_id).Find(&agentModels).Error; err != nil {
		return nil, err
	}
	res := make([]*entity.Agent, 0, len(agentModels))
	for _, agentModel := range agentModels {
		res = append(res, agentModel.Domain())
	}
	return res, nil
}
