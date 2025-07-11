package gorminterface

import (
	"context"

	"github.com/wang900115/Perry/internal/adapter/validator"
	"github.com/wang900115/Perry/internal/domain/entity"
)

type ToDo interface {
	// 新增任務事項
	Create(context.Context, uint, validator.ToDoCreateRequest) (*entity.ToDo, error)
	// 更新任務事項
	Update(context.Context, validator.ToDoUpdateRequest) (*entity.ToDo, error)
	// 刪除任務事項
	Delete(context.Context, validator.ToDoDeleteRequest) error
	// 取得任務事項
	Query(context.Context, uint) ([]*entity.ToDo, error)
	// 取得任務狀態
	GetStatus(context.Context, validator.ToDoGetStatusRequest) (string, error)
}
