package gorminterface

import (
	"context"

	"github.com/wang900115/Perry/internal/adapter/validator"
	"github.com/wang900115/Perry/internal/domain/entity"
)

type Agent interface {
	// 增加一個代理身分
	Add(context.Context, uint, validator.AgentAddRequest) (*entity.Agent, error)
	// 扣除一個代理身分
	Remove(context.Context, validator.AgentRemoveRequest) error
	// 扣除該用戶全部的代理身分
	RemoveAll(context.Context, uint) error
	// 取得使用者下全部的代理身分
	Read(context.Context, uint) ([]*entity.Agent, error)
}
