package redistable

import (
	"time"

	rediskey "github.com/wang900115/Perry/internal/adapter/redis/key"
	"github.com/wang900115/Perry/internal/domain/entity"
	"github.com/wang900115/utils/convert"
)

type ToDo struct {
	Name      string    `json:"name"`
	Priority  string    `json:"priority"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Status    string    `json:"status"`
}

// redis's model -> redis
func (t ToDo) ToHash() map[string]interface{} {
	return map[string]interface{}{
		rediskey.REDIS_FIELD_TODO_NAME:      t.Name,
		rediskey.REDIS_FIELD_TODO_PRIORITY:  t.Priority,
		rediskey.REDIS_FIELD_TODO_STARTTIME: t.StartTime,
		rediskey.REDIS_FIELD_TODO_ENDTIME:   t.EndTime,
		rediskey.REDIS_FIELD_TODO_STATUS:    t.Status,
	}
}

// !TODO(utils) redis's model <- redis
func (t ToDo) FromHash(data map[string]string) ToDo {
	return ToDo{
		Name:      data[rediskey.REDIS_FIELD_TODO_NAME],
		Priority:  data[rediskey.REDIS_FIELD_TODO_PRIORITY],
		StartTime: convert.FromStringToTimeTime(data[rediskey.REDIS_FIELD_TODO_STARTTIME]),
		EndTime:   convert.FromStringToTimeTime(data[rediskey.REDIS_FIELD_TODO_ENDTIME]),
		Status:    data[rediskey.REDIS_FIELD_TODO_STATUS],
	}
}

func (t ToDo) ToDomain() *entity.ToDo {
	return &entity.ToDo{
		Name:      t.Name,
		Priority:  t.Priority,
		StartTime: t.StartTime,
		EndTime:   t.EndTime,
		Status:    t.Status,
	}
}
