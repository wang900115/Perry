package redisinterface

import (
	"context"

	redistable "github.com/wang900115/Perry/internal/adapter/redis/table"
)

type Agent interface {
	// 初始化 (key 為 agent_id 預設為正常狀態)
	Initialize(context.Context, uint) error
	// 更新
	Update(context.Context, uint, string) error
	// 取得使用者底下的代理
	Get(context.Context, uint) ([]*redistable.Agent, error)
	// 刪除特定代理
	Delete(context.Context, uint) error
	// 刪除使用者底下的全部代理 (user_id)
	DeleteAll(context.Context, uint) error
}
