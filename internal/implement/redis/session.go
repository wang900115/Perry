package redisimplement

import (
	"context"

	"github.com/redis/go-redis/v9"
	rediskey "github.com/wang900115/Perry/internal/adapter/redis/key"
	redistable "github.com/wang900115/Perry/internal/adapter/redis/table"
	redisinterface "github.com/wang900115/Perry/internal/domain/interface/redis"
	"github.com/wang900115/utils/convert"
)

type Session struct {
	redis *redis.Client
}

func NewSessionImplement(redis *redis.Client) redisinterface.Session {
	return &Session{redis: redis}
}

func (s *Session) Generate(ctx context.Context, userId uint, ip, userAgent string) (int64, error) {
	sessionID, err := s.redis.Incr(ctx, rediskey.REDIS_INCR_USER_SESSION_ID).Result()
	if err != nil {
		return 0, err
	}
	session := redistable.UserSession{
		IP:        ip,
		UserAgent: userAgent,
		Provider:  "local",
		IsActive:  true,
	}
	sessionId := convert.FromInt64ToString(sessionID)
	key := rediskey.REDIS_TABLE_USER_SESSION + string(sessionId)
	err = s.redis.HSet(ctx, key, session.ToHash()).Err()
	if err != nil {
		return 0, err
	}
	listKey := rediskey.REDIS_LIST_USER_SESSION + convert.FromUintToString(userId)
	err = s.redis.LPush(ctx, listKey, sessionId).Err()
	if err != nil {
		return 0, err
	}
	err = s.redis.LTrim(ctx, listKey, 0, 2).Err()
	if err != nil {
		return 0, err
	}
	return sessionID, nil
}

func (s *Session) Get(ctx context.Context, sessionId int64) (*redistable.UserSession, error) {
	key := rediskey.REDIS_TABLE_USER_SESSION + convert.FromInt64ToString(sessionId)
	data, err := s.redis.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return redistable.UserSession{}.FromHash(data), nil
}

func (s *Session) Deactivate(ctx context.Context, sessionId int64) error {
	key := rediskey.REDIS_TABLE_USER_SESSION + convert.FromInt64ToString(sessionId)
	field := rediskey.REDIS_FIELD_USER_SESSION_IS_ACTIVE
	if err := s.redis.HSet(ctx, key, field, "0").Err(); err != nil {
		return err
	}
	return nil
}

func (s *Session) Delete(ctx context.Context, sessionId int64) error {
	key := rediskey.REDIS_TABLE_USER_SESSION + convert.FromInt64ToString(sessionId)
	if err := s.redis.Del(ctx, key).Err(); err != nil {
		return err
	}
	return nil
}
