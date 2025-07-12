package gormimplement

import (
	"context"

	gormmodel "github.com/wang900115/Perry/internal/adapter/gorm/model"
	"github.com/wang900115/Perry/internal/adapter/validator"
	"github.com/wang900115/Perry/internal/domain/entity"
	gorminterface "github.com/wang900115/Perry/internal/domain/interface/gorm"
	"gorm.io/gorm"
)

type ToDo struct {
	gorm *gorm.DB
}

func NewToDoImplement(gorm *gorm.DB) gorminterface.ToDo {
	return &ToDo{gorm: gorm}
}

func (t *ToDo) Create(ctx context.Context, user_id uint, input validator.ToDoCreateRequest) (*entity.ToDo, error) {
	createdToDo := gormmodel.ToDo{
		UserID:    user_id,
		AgentID:   input.ID,
		Name:      input.Name,
		Priority:  input.Priority,
		StartTime: input.StartTime,
		EndTime:   input.EndTime,
	}
	if err := t.gorm.WithContext(ctx).Create(&createdToDo).Error; err != nil {
		return nil, err
	}

	return createdToDo.ToDomain(), nil
}

func (t *ToDo) Update(ctx context.Context, input validator.ToDoUpdateRequest) (*entity.ToDo, error) {
	if err := t.gorm.WithContext(ctx).Model(&gormmodel.ToDo{}).Where("id = ?", input.ID).Updates(map[string]interface{}{
		"name":       input.Name,
		"priority":   input.Priority,
		"start_time": input.StartTime,
		"end_time":   input.EndTime,
		"status":     input.Status,
	}).Error; err != nil {
		return nil, err
	}

	var todoModel gormmodel.ToDo
	if err := t.gorm.WithContext(ctx).First(&todoModel, input.ID).Error; err != nil {
		return nil, err
	}

	return todoModel.ToDomain(), nil
}

func (t *ToDo) Delete(ctx context.Context, input validator.ToDoDeleteRequest) error {
	var todoModel gormmodel.ToDo
	if err := t.gorm.WithContext(ctx).Delete(&todoModel, input.ID).Error; err != nil {
		return err
	}
	return nil
}

func (t *ToDo) Query(ctx context.Context, user_id uint) ([]*entity.ToDo, error) {
	var todoModels []gormmodel.ToDo
	if err := t.gorm.WithContext(ctx).Where("user_id = ?", user_id).Find(&todoModels).Error; err != nil {
		return nil, err
	}
	res := make([]*entity.ToDo, 0, len(todoModels))
	for _, todoModel := range todoModels {
		res = append(res, todoModel.ToDomain())
	}
	return res, nil
}

func (t *ToDo) QueryAgent(ctx context.Context, agent_id uint) ([]*entity.ToDo, error) {
	var todoModels []gormmodel.ToDo
	if err := t.gorm.WithContext(ctx).Where("agent_id = ?", agent_id).Find(&todoModels).Error; err != nil {
		return nil, err
	}
	res := make([]*entity.ToDo, 0, len(todoModels))
	for _, todoModel := range todoModels {
		res = append(res, todoModel.ToDomain())
	}
	return res, nil
}
