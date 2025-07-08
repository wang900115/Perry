package redisimplement

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/redis/go-redis/v9"
	rediskey "github.com/wang900115/Perry/internal/adapter/redis/key"
	redistable "github.com/wang900115/Perry/internal/adapter/redis/table"
	"github.com/wang900115/utils/convert"
)

type Token struct {
	redis      *redis.Client
	expiration int64
	issuer     string
	secret     []byte
}

func NewTokenImplement(redis *redis.Client, expiration int64, issuer, secret string) *Token {
	return &Token{redis: redis, expiration: expiration, issuer: issuer, secret: []byte(secret)}
}

func (t *Token) Generate(ctx context.Context, userID uint, sessionID int64) (string, error) {
	now := time.Now().Unix()
	exp := now + t.expiration
	claims := redistable.Claims{
		StandardClaims: jwt.StandardClaims{
			Id:        convert.FromInt64ToString(sessionID),
			Subject:   convert.FromUintToString(userID),
			IssuedAt:  now,
			ExpiresAt: exp,
			Issuer:    t.issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256,
		claims.StandardClaims)
	signedToken, err := token.SignedString(t.secret)
	if err != nil {
		return "", err
	}
	key := groupName(userID, sessionID)
	if err := t.redis.HSet(ctx, key, claims.ToHash()).Err(); err != nil {
		return "", err
	}
	if err := t.redis.ExpireAt(ctx, key, time.Unix(exp, 0)).Err(); err != nil {
		return "", err
	}
	return signedToken, nil
}

func (t *Token) Validate(ctx context.Context, token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return t.secret, nil
	})
	if err != nil {
		return err
	}

	claims, ok := parsedToken.Claims.(*jwt.StandardClaims)
	if !ok || !parsedToken.Valid {
		return errors.New("invalid claims or token")
	}

	now := time.Now().Unix()
	if claims.ExpiresAt > 0 && claims.ExpiresAt < now {
		return errors.New("token is expired")
	}

	key := groupNameString(claims.Subject, claims.Id)
	exists, err := t.redis.Exists(ctx, key).Result()
	if err != nil {
		return err
	}
	if exists == 0 {
		return errors.New("token is not found in redis")
	}
	return nil
}

func (t *Token) Refresh(ctx context.Context, token string) (string, error) {
	oldToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return t.secret, nil
	})
	if err != nil {
		return "", err
	}

	claims := oldToken.Claims.(*jwt.StandardClaims)
	sessionID, err := strconv.ParseInt(claims.Id, 10, 64)
	if err != nil {
		return "", err
	}
	userID, err := strconv.ParseUint(claims.Subject, 10, 64)
	if err != nil {
		return "", err
	}
	return t.Generate(ctx, uint(userID), sessionID)
}

func (t *Token) Delete(ctx context.Context, userID uint, sessionID int64) error {
	return t.redis.Del(ctx, groupName(userID, sessionID)).Err()
}

func (t *Token) DeleteAll(ctx context.Context, userID uint) error {
	pattern := Name(userID)

	iter := t.redis.Scan(ctx, 0, pattern, 100).Iterator()
	var keysToDelete []string

	for iter.Next(ctx) {
		keysToDelete = append(keysToDelete, iter.Val())

		if len(keysToDelete) >= 100 {
			if err := t.redis.Del(ctx, keysToDelete...).Err(); err != nil {
				return err
			}
			keysToDelete = keysToDelete[:0]
		}
	}

	if len(keysToDelete) > 0 {
		if err := t.redis.Del(ctx, keysToDelete...).Err(); err != nil {
			return err
		}
	}

	if err := iter.Err(); err != nil {
		return err
	}

	return nil
}

func Name(userID uint) string {
	return fmt.Sprintf("%s:%d", rediskey.REDIS_TABLE_USER_TOKEN, userID)
}

func groupName(userID uint, sessionID int64) string {
	return fmt.Sprintf("%s:%d:%d", rediskey.REDIS_TABLE_USER_TOKEN, userID, sessionID)
}

func groupNameString(userID, sessionID string) string {
	return fmt.Sprintf("%s:%s:%s", rediskey.REDIS_TABLE_USER_TOKEN, userID, sessionID)
}
