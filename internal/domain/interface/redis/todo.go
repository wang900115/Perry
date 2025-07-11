package redisinterface

import (
	"context"

	redistable "github.com/wang900115/Perry/internal/adapter/redis/table"
	"github.com/wang900115/Perry/internal/adapter/validator"
)

type ToDo interface {
	// 初始化 (key 為 todo_id 預設為執行狀態)
	Initialize(context.Context, uint) error
	// 更新任務資訊
	Update(context.Context, validator.ToDoUpdateRequest) (redistable.ToDo, error)
	// 取得使用者下任務資訊
	Get(context.Context, uint) ([]*redistable.ToDo, error)
	// 刪除任務
	Delete(context.Context, validator.ToDoDeleteRequest) error
	// 取得代理者下的任務
	GetAgent(context.Context, uint) ([]*redistable.ToDo, error)
}
