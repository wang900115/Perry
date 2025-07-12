package redisinterface

import (
	"context"

	"github.com/wang900115/Perry/internal/adapter/validator"
	"github.com/wang900115/Perry/internal/domain/entity"
)

type Agent interface {
	// 初始化 (key 為 agent_id )
	Initialize(context.Context, uint, validator.AgentAddRequest) error
	// 取得使用者底下的代理
	Get(context.Context, uint) ([]*entity.Agent, error)
	// 刪除特定代理
	Delete(context.Context, uint) error
	// 刪除使用者底下的全部代理 (user_id)
	DeleteAll(context.Context, uint) error
}
