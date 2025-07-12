package redisimplement

import (
	"context"

	"github.com/redis/go-redis/v9"
	rediskey "github.com/wang900115/Perry/internal/adapter/redis/key"
	redistable "github.com/wang900115/Perry/internal/adapter/redis/table"
	"github.com/wang900115/Perry/internal/adapter/validator"
	"github.com/wang900115/Perry/internal/domain/entity"
	redisinterface "github.com/wang900115/Perry/internal/domain/interface/redis"
	"github.com/wang900115/utils/convert"
)

type Agent struct {
	redis *redis.Client
}

func NewAgentImplement(redis *redis.Client) redisinterface.Agent {
	return &Agent{redis: redis}
}

func (a *Agent) Initialize(ctx context.Context, agent_id uint, input validator.AgentAddRequest) error {
	key := rediskey.REDIS_TABLE_AGENT + convert.FromUintToString(agent_id)

	agent := redistable.Agent{
		Name:        input.Name,
		Age:         input.Age,
		Role:        input.Role,
		Language:    input.Language,
		Description: input.Description,
		Status:      input.Status,
	}

	err := a.redis.HSet(ctx, key, agent.ToHash()).Err()
	if err != nil {
		return err
	}
	return nil
}

func (a *Agent) Get(ctx context.Context, user_id uint) ([]*entity.Agent, error) {
	key := rediskey.REDIS_LIST_USER_AGENT + convert.FromUintToString(user_id)

	agentIDs, err := a.redis.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var agents []*entity.Agent
	for _, id := range agentIDs {
		key := rediskey.REDIS_TABLE_AGENT + id

		data, err := a.redis.HGetAll(ctx, key).Result()
		if err != nil {
			return nil, err
		}

		agents = append(agents, redistable.Agent{}.FromHash(data).ToDomain())
	}

	return agents, nil
}

func (a *Agent) Delete(ctx context.Context, agent_id uint) error {
	key := rediskey.REDIS_TABLE_AGENT + convert.FromUintToString(agent_id)
	if err := a.redis.Del(ctx, key).Err(); err != nil {
		return err
	}
	return nil
}

func (a *Agent) DeleteAll(ctx context.Context, user_id uint) error {
	key := rediskey.REDIS_LIST_USER_AGENT + convert.FromUintToString(user_id)

	agentIDs, err := a.redis.HGetAll(ctx, key).Result()
	if err != nil {
		return err
	}

	for _, id := range agentIDs {
		key := rediskey.REDIS_TABLE_AGENT + id

		err := a.redis.Del(ctx, key).Err()
		if err != nil {
			return err
		}

	}

	if err := a.redis.Del(ctx, key).Err(); err != nil {
		return err
	}

	return nil
}
