package redistable

import (
	rediskey "github.com/wang900115/Perry/internal/adapter/redis/key"
	"github.com/wang900115/Perry/internal/domain/entity"
)

type Agent struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Role        string `json:"role"`
	Language    string `json:"language"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// redis's model -> redis
func (a Agent) ToHash() map[string]interface{} {
	return map[string]interface{}{
		rediskey.REDIS_FIELD_AGENT_NAME:        a.Name,
		rediskey.REDIS_FIELD_AGENT_AGE:         a.Age,
		rediskey.REDIS_FIELD_AGENT_ROLE:        a.Role,
		rediskey.REDIS_FIELD_AGENT_LANGUAGE:    a.Language,
		rediskey.REDIS_FIELD_AGENT_DESCRIPTION: a.Description,
		rediskey.REDIS_FIELD_AGENT_STATUS:      a.Status,
	}
}

// !todo redis's model <- redis
func (a Agent) FromHash(data map[string]string) Agent {
	return Agent{
		Name:        data[rediskey.REDIS_FIELD_AGENT_NAME],
		Age:         data[rediskey.REDIS_FIELD_AGENT_AGE],
		Role:        data[rediskey.REDIS_FIELD_AGENT_ROLE],
		Language:    data[rediskey.REDIS_FIELD_AGENT_LANGUAGE],
		Description: data[rediskey.REDIS_FIELD_AGENT_DESCRIPTION],
		Status:      data[rediskey.REDIS_FIELD_AGENT_STATUS],
	}
}

// redis's model -> Domain
func (a Agent) ToDomain() *entity.Agent {
	return &entity.Agent{
		Name:        a.Name,
		Age:         a.Age,
		Role:        a.Role,
		Language:    a.Language,
		Description: a.Description,
		Status:      a.Status,
	}
}
