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
