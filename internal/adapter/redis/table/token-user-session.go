package redistable

import (
	"github.com/golang-jwt/jwt"
	rediskey "github.com/wang900115/Perry/internal/adapter/redis/key"
	"github.com/wang900115/utils/convert"
)

type Claims struct {
	UserID         uint
	StandardClaims jwt.StandardClaims
}

// redis's model -> redis
func (c Claims) ToHash() map[string]interface{} {
	return map[string]interface{}{
		rediskey.REDIS_FIELD_USER_TOKEN_USER_ID: convert.FromUintToString(c.UserID),
		// rediskey.REDIS_FIELD_USER_TOKEN_AUD:      c.Audience,
		rediskey.REDIS_FIELD_USER_TOKEN_EXP:      convert.FromTimestampToString(c.StandardClaims.ExpiresAt),
		rediskey.REDIS_FIELD_USER_TOKEN_ISSUEDAT: convert.FromTimestampToString(c.StandardClaims.IssuedAt),
		rediskey.REDIS_FIELD_USER_TOKEN_ISSUER:   c.StandardClaims.Issuer,
		rediskey.REDIS_FIELD_USER_TOKEN_JTI:      c.StandardClaims.Id,
		// rediskey.REDIS_FIELD_USER_TOKEN_NBF:      convert.FromTimestampToString(c.NotBefore),
		rediskey.REDIS_FIELD_USER_TOKEN_SUB: c.StandardClaims.Subject,
	}
}

// redis's model <- redis
func (c Claims) FromHash(data map[string]string) Claims {
	return Claims{
		UserID: convert.FromStringToUint(data[rediskey.REDIS_FIELD_USER_TOKEN_USER_ID]),

		StandardClaims: jwt.StandardClaims{
			// Audience:  data[rediskey.REDIS_FIELD_USER_TOKEN_AUD],
			ExpiresAt: convert.FromStringToTimeStamp(data[rediskey.REDIS_FIELD_USER_TOKEN_EXP]),
			Id:        data[rediskey.REDIS_FIELD_USER_TOKEN_JTI],
			IssuedAt:  convert.FromStringToTimeStamp(data[rediskey.REDIS_FIELD_USER_TOKEN_ISSUEDAT]),
			Issuer:    data[rediskey.REDIS_FIELD_USER_TOKEN_ISSUER],
			// NotBefore: convert.FromStringToTimeStamp(data[rediskey.REDIS_FIELD_USER_TOKEN_NBF]),
			Subject: data[rediskey.REDIS_FIELD_USER_TOKEN_SUB],
		},
	}
}
