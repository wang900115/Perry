package redistable

import rediskey "github.com/wang900115/Perry/internal/adapter/redis/key"

type ToDo struct {
	Status string `json:"status"`
}

// redis's model -> redis
func (t ToDo) ToHash() map[string]interface{} {
	return map[string]interface{}{
		rediskey.REDIS_FIELD_TODO_STATUS: t.Status,
	}
}

// redis's model <- redis
func (t ToDo) FromHash(data map[string]string) ToDo {
	return ToDo{
		Status: data[rediskey.REDIS_FIELD_TODO_STATUS],
	}
}
