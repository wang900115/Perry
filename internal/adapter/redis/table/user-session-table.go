package redistable

import (
	rediskey "github.com/wang900115/Perry/internal/adapter/redis/key"
	"github.com/wang900115/utils/convert"
)

type UserSession struct {
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	Provider  string `json:"provider"`
	CreatedAt int64  `json:"created_at"`
	ExpiredAt int64  `json:"expired_at"`
	IsActive  bool   `json:"active"`
}

// redis's model -> redis
func (u UserSession) ToHash() map[string]interface{} {
	return map[string]interface{}{
		rediskey.REDIS_FIELD_USER_SESSION_IP:        u.IP,
		rediskey.REDIS_FIELD_USER_SESSION_PROVIDER:  u.Provider,
		rediskey.REDIS_FIELD_USER_SESSION_UA:        u.UserAgent,
		rediskey.REDIS_FIELD_USER_SESSION_CREATED:   convert.FromTimestampToString(u.CreatedAt),
		rediskey.REDIS_FIELD_USER_SESSION_EXPIRED:   convert.FromTimestampToString(u.ExpiredAt),
		rediskey.REDIS_FIELD_USER_SESSION_IS_ACTIVE: convert.FromBoolToString(u.IsActive),
	}
}

// redis's model <- redis
func (u UserSession) FromHash(data map[string]string) *UserSession {
	return &UserSession{
		Provider:  data[rediskey.REDIS_FIELD_USER_SESSION_PROVIDER],
		IP:        data[rediskey.REDIS_FIELD_USER_SESSION_IP],
		UserAgent: data[rediskey.REDIS_FIELD_USER_SESSION_UA],
		CreatedAt: convert.FromStringToTimeStamp(data[rediskey.REDIS_FIELD_USER_SESSION_CREATED]),
		ExpiredAt: convert.FromStringToTimeStamp(data[rediskey.REDIS_FIELD_USER_SESSION_EXPIRED]),
		IsActive:  convert.FromStringToBool(data[rediskey.REDIS_FIELD_USER_SESSION_IS_ACTIVE]),
	}
}
