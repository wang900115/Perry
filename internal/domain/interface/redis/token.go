package redisinterface

import (
	"context"
)

type Token interface {
	// 產生 token (key 為 userId sessionId)
	Generate(context.Context, uint, int64) (string, error)
	// 驗證 token
	Validate(context.Context, string) error
	// 重發 token
	Refresh(context.Context, string) (string, error)
	// 刪除特定 token
	Delete(context.Context, uint, int64) error
	// 刪除全部 token 包含 session
	DeleteAll(context.Context, uint) error
}
