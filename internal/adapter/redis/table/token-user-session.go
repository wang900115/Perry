package redistable

import (
	"github.com/golang-jwt/jwt"
	rediskey "github.com/wang900115/Perry/internal/adapter/redis/key"
	"github.com/wang900115/utils/convert"
)

type Claims struct {
	jwt.StandardClaims
}

// redis's model -> redis
func (c Claims) ToHash() map[string]interface{} {
	return map[string]interface{}{
		// rediskey.REDIS_FIELD_USER_TOKEN_AUD:      c.Audience,
		rediskey.REDIS_FIELD_USER_TOKEN_EXP:      convert.FromTimestampToString(c.ExpiresAt),
		rediskey.REDIS_FIELD_USER_TOKEN_ISSUEDAT: convert.FromTimestampToString(c.IssuedAt),
		rediskey.REDIS_FIELD_USER_TOKEN_ISSUER:   c.Issuer,
		rediskey.REDIS_FIELD_USER_TOKEN_JTI:      c.Id,
		// rediskey.REDIS_FIELD_USER_TOKEN_NBF:      convert.FromTimestampToString(c.NotBefore),
		rediskey.REDIS_FIELD_USER_TOKEN_SUB: c.Subject,
	}
}

// redis's model <- redis
func (c Claims) FromHash(data map[string]string) Claims {
	return Claims{

		jwt.StandardClaims{
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
