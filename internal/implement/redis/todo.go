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

type ToDo struct {
	redis *redis.Client
}

func NewToDoImplement(redis *redis.Client) redisinterface.ToDo {
	return &ToDo{redis: redis}
}

func (t *ToDo) Initialize(ctx context.Context, todo_id uint, input validator.ToDoCreateRequest) error {
	key := rediskey.REDIS_TABLE_TODO + convert.FromUintToString(todo_id)
	todo := redistable.ToDo{
		Name:      input.Name,
		Priority:  input.Priority,
		StartTime: input.StartTime,
		EndTime:   input.EndTime,
		Status:    input.Status,
	}
	err := t.redis.HSet(ctx, key, todo.ToHash()).Err()
	if err != nil {
		return err
	}

	agentKey := rediskey.REDIS_LIST_AGENT_TODO + convert.FromUintToString(input.ID)
	if err := t.redis.RPush(ctx, agentKey, todo_id).Err(); err != nil {
		return err
	}
	return nil
}

func (t *ToDo) Update(ctx context.Context, input validator.ToDoUpdateRequest) (*entity.ToDo, error) {
	key := rediskey.REDIS_TABLE_TODO + convert.FromUintToString(input.ID)
	todo := redistable.ToDo{
		Name:      input.Name,
		Priority:  input.Priority,
		StartTime: input.StartTime,
		EndTime:   input.EndTime,
		Status:    input.Status,
	}
	err := t.redis.HSet(ctx, key, todo.ToHash()).Err()
	if err != nil {
		return nil, err
	}
	return todo.ToDomain(), nil
}

func (t *ToDo) GetUser(ctx context.Context, user_id uint) ([]*entity.ToDo, error) {
	key := rediskey.REDIS_LIST_USER_TODO + convert.FromUintToString(user_id)
	todoIDs, err := t.redis.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var todos []*entity.ToDo
	for _, id := range todoIDs {
		key := rediskey.REDIS_TABLE_TODO + id

		data, err := t.redis.HGetAll(ctx, key).Result()
		if err != nil {
			return nil, err
		}

		todos = append(todos, redistable.ToDo{}.FromHash(data).ToDomain())
	}
	return todos, nil
}

func (t *ToDo) GetAgent(ctx context.Context, agent_id uint) ([]*entity.ToDo, error) {
	key := rediskey.REDIS_LIST_AGENT_TODO + convert.FromUintToString(agent_id)
	todoIDs, err := t.redis.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var todos []*entity.ToDo
	for _, id := range todoIDs {
		key := rediskey.REDIS_TABLE_TODO + id

		data, err := t.redis.HGetAll(ctx, key).Result()
		if err != nil {
			return nil, err
		}

		todos = append(todos, redistable.ToDo{}.FromHash(data).ToDomain())
	}
	return todos, nil
}

func (t *ToDo) Delete(ctx context.Context, todo_id uint) error {
	key := rediskey.REDIS_TABLE_TODO + convert.FromUintToString(todo_id)
	if err := t.redis.Del(ctx, key).Err(); err != nil {
		return err
	}
	return nil
}

func (t *ToDo) DeleteUser(ctx context.Context, user_id uint) error {
	key := rediskey.REDIS_LIST_USER_TODO + convert.FromUintToString(user_id)
	todoIDs, err := t.redis.HGetAll(ctx, key).Result()
	if err != nil {
		return err
	}
	for _, id := range todoIDs {
		key := rediskey.REDIS_TABLE_TODO + id
		err := t.redis.Del(ctx, key).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *ToDo) DeleteAgent(ctx context.Context, agent_id uint) error {
	key := rediskey.REDIS_LIST_AGENT_TODO + convert.FromUintToString(agent_id)
	todoIDs, err := t.redis.HGetAll(ctx, key).Result()
	if err != nil {
		return err
	}
	for _, id := range todoIDs {
		key := rediskey.REDIS_TABLE_TODO + id
		err := t.redis.Del(ctx, key).Err()
		if err != nil {
			return err
		}
	}
	return nil
}
