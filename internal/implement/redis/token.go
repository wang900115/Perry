package redisimplement

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Token struct {
	redis *redis.Client
}

func NewTokenImplement(redis *redis.Client) *Token {
	return &Token{redis: redis}
}

func (t *Token) Generate(ctx context.Context, userID uint, sessionID int64) (string, error) {

}

func (t *Token) Validate(ctx context.Context, token string) error {

}

func (t *Token) Refresh(ctx context.Context, token string) (string, error) {

}

func (t *Token) Delete(ctx context.Context, userID uint, sessionID int64) error {

}

func (t *Token) DeleteAll(ctx context.Context, userID uint) error {

}
