package redisinterface

import (
	"context"
)

type Agent interface {
	// 初始化 (key 為 agent_id 預設為正常狀態)
	Initialize(context.Context, uint) error
	// 更新狀態
	UpdateStatus(context.Context, uint, string) error
	// 取得狀態
	GetStatus(context.Context, uint) (string, error)
}
