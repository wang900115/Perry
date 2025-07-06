package redisimplement

import (
	"context"

	"github.com/redis/go-redis/v9"
	redistable "github.com/wang900115/Perry/internal/adapter/redis/table"
	redisinterface "github.com/wang900115/Perry/internal/domain/interface/redis"
)

type Session struct {
	redis *redis.Client
}

func NewSessionImplement(redis *redis.Client) redisinterface.Session {
	return &Session{redis: redis}
}

func (s *Session) Generate(ctx context.Context, userId uint, ip, userAgent string) (int64, error) {

}

func (s *Session) Get(ctx context.Context, sessionId int64) (redistable.UserSession, error) {

}

func (s *Session) Deactivate(ctx context.Context, sessionId int64) error {

}

func (s *Session) Delete(ctx context.Context, sessionId int64) error {

}
