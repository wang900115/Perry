package redistable

import rediskey "github.com/wang900115/Perry/internal/adapter/redis/key"

type Agent struct {
	Status string `json:"status"`
}

// redis's model -> redis
func (a Agent) ToHash() map[string]interface{} {
	return map[string]interface{}{
		rediskey.REDIS_FIELD_AGENT_STATUS: a.Status,
	}
}

// redis's model <- redis
func (a Agent) FromHash(data map[string]string) Agent {
	return Agent{
		Status: data[rediskey.REDIS_FIELD_AGENT_STATUS],
	}
}
