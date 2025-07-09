package rediskey

const (
	REDIS_TABLE_USER_STATUS  = "user:status"   // key 為 user_id
	REDIS_TABLE_USER_SESSION = "user:session:" // key 為 session_id (最多只有3個)
	REDIS_TABLE_USER_TOKEN   = "user:jwt:"     // key 為 token string
)

const (
	REDIS_FIELD_USER_STATUS_USERNAME   = "username"
	REDIS_FIELD_USER_STATUS_FULLNAME   = "fullname"
	REDIS_FIELD_USER_STATUS_NICKNAME   = "nickname"
	REDIS_FIELD_USER_STATUS_AVATARURL  = "avatarurl"
	REDIS_FIELD_USER_STATUS_DEVICE     = "device"
	REDIS_FIELD_USER_STATUS_LASTIP     = "last_ip"
	REDIS_FIELD_USER_STATUS_LASTLOGIN  = "last_login"
	REDIS_FIELD_USER_STATUS_LASTLOGOUT = "last_logout"
)

// session
const (
	REDIS_INCR_USER_SESSION_ID = "session_id"
)
const (
	REDIS_FIELD_USER_SESSION_PROVIDER  = "provider"
	REDIS_FIELD_USER_SESSION_IP        = "ip"
	REDIS_FIELD_USER_SESSION_UA        = "user_agent"
	REDIS_FIELD_USER_SESSION_CREATED   = "created_at"
	REDIS_FIELD_USER_SESSION_EXPIRED   = "expired_at"
	REDIS_FIELD_USER_SESSION_IS_ACTIVE = "is_active"
)

// token
const (
	REDIS_FIELD_USER_TOKEN_USER_ID = "user_id"

	REDIS_FIELD_USER_TOKEN_AUD      = "aud"
	REDIS_FIELD_USER_TOKEN_JTI      = "jti"
	REDIS_FIELD_USER_TOKEN_ISSUER   = "issuer"
	REDIS_FIELD_USER_TOKEN_ISSUEDAT = "issued_at"
	REDIS_FIELD_USER_TOKEN_EXP      = "expired_at"
	REDIS_FIELD_USER_TOKEN_NBF      = "not_before"
	REDIS_FIELD_USER_TOKEN_SUB      = "subject"
)
