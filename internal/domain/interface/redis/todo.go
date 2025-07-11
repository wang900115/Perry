package redisinterface

import (
	"context"

	"github.com/wang900115/Perry/internal/adapter/validator"
)

type ToDo interface {
	// 初始化 (ley 為 todo_id 預設為執行狀態)
	Initialize(context.Context, uint) error
	// 更新狀態
	UpdateStatus(context.Context, uint, string) error
	// 取得狀態
	GetStatus(context.Context, validator.ToDoGetStatusRequest) (string, error)
	// 刪除任務
	Delete(context.Context, validator.ToDoDeleteRequest) error
}
