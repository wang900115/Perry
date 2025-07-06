package redistable

import (
	rediskey "github.com/wang900115/Perry/internal/adapter/redis/key"
	"github.com/wang900115/Perry/internal/domain/entity"
	"github.com/wang900115/utils/convert"
)

type UserStatus struct {
	Username  string `json:"username"`
	FullName  string `json:"fullname"`
	NickName  string `json:"nickname"`
	AvatarURL string `json:"avatarURL"`

	Device     string `json:"device"`
	LastIP     string `json:"last_ip"`
	LastLogin  int64  `json:"last_login"`
	LastLogout int64  `json:"last_logout"`
}

// redis's model -> redis
func (u UserStatus) ToHash() map[string]interface{} {
	return map[string]interface{}{
		rediskey.REDIS_FIELD_USER_STATUS_USERNAME:  u.Username,
		rediskey.REDIS_FIELD_USER_STATUS_FULLNAME:  u.FullName,
		rediskey.REDIS_FIELD_USER_STATUS_NICKNAME:  u.NickName,
		rediskey.REDIS_FIELD_USER_STATUS_AVATARURL: u.AvatarURL,

		rediskey.REDIS_FIELD_USER_STATUS_DEVICE:     u.Device,
		rediskey.REDIS_FIELD_USER_STATUS_LASTIP:     u.LastIP,
		rediskey.REDIS_FIELD_USER_STATUS_LASTLOGIN:  u.LastLogin,
		rediskey.REDIS_FIELD_USER_STATUS_LASTLOGOUT: u.LastLogout,
	}
}

// redis's model <- redis
func (u UserStatus) FromHash(data map[string]string) UserStatus {
	return UserStatus{
		Username:  data[rediskey.REDIS_FIELD_USER_STATUS_USERNAME],
		FullName:  data[rediskey.REDIS_FIELD_USER_STATUS_FULLNAME],
		NickName:  data[rediskey.REDIS_FIELD_USER_STATUS_NICKNAME],
		AvatarURL: data[rediskey.REDIS_FIELD_USER_STATUS_AVATARURL],

		Device:     data[rediskey.REDIS_FIELD_USER_STATUS_DEVICE],
		LastIP:     data[rediskey.REDIS_FIELD_USER_STATUS_LASTIP],
		LastLogin:  convert.FromStringToTimeStamp(data[rediskey.REDIS_FIELD_USER_STATUS_LASTLOGIN]),
		LastLogout: convert.FromStringToTimeStamp(data[rediskey.REDIS_FIELD_USER_STATUS_LASTLOGOUT]),
	}
}

// redis's model -> domain
func (u UserStatus) ToDomain() entity.UserStatus {
	return entity.UserStatus{
		User: entity.User{
			Username:  u.Username,
			FullName:  u.FullName,
			NickName:  u.NickName,
			AvatarURL: u.AvatarURL,
		},

		Device:     u.Device,
		LastIP:     u.LastIP,
		LastLogin:  u.LastLogin,
		LastLogout: u.LastLogout,
	}
}
