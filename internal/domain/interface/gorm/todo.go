package gorminterface

import (
	"context"

	"github.com/wang900115/Perry/internal/adapter/validator"
	"github.com/wang900115/Perry/internal/domain/entity"
)

type ToDo interface {
	// 新增待辦事項
	Create(context.Context, uint, validator.ToDoCreateRequest) (*entity.ToDo, error)
	// 更新待辦事項
	Update(context.Context, uint, validator.ToDoUpdateRequest) (*entity.ToDo, error)
	// 刪除待辦事項
	Delete(context.Context, uint, validator.ToDoDeleteRequest) error
	// 取得待辦事項
	Query(context.Context, uint) ([]*entity.ToDo, error)
}
