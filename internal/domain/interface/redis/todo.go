package redisinterface

import (
	"context"

	"github.com/wang900115/Perry/internal/adapter/validator"
	"github.com/wang900115/Perry/internal/domain/entity"
)

type ToDo interface {
	// 創建任務並分發
	Initialize(context.Context, uint, validator.ToDoCreateRequest) error
	// 更新任務資訊
	Update(context.Context, validator.ToDoUpdateRequest) (*entity.ToDo, error)
	// 取得使用者下任務資訊
	GetUser(context.Context, uint) ([]*entity.ToDo, error)
	// 取得代理者下的任務
	GetAgent(context.Context, uint) ([]*entity.ToDo, error)
	// 刪除任務
	Delete(context.Context, uint) error
	// 刪除用戶下的任務
	DeleteUser(context.Context, uint) error
	// 刪除代理下的任務
	DeleteAgent(context.Context, uint) error
}
