package redisinterface

import (
	"context"

	redistable "github.com/wang900115/Perry/internal/adapter/redis/table"
)

type Session interface {
	//  創建 session(key為sessionid)
	Generate(context.Context, uint, string, string) (int64, error)
	//  拿取 session(依據sessionid)
	Get(context.Context, int64) (redistable.UserSession, error)
	//  禁止 session(依據sessionid)
	Deactivate(context.Context, int64) error
	//  刪除 session(依據sessionid)
	Delete(context.Context, int64) error
}
