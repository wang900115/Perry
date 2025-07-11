package usecase

import (
	"context"

	"github.com/wang900115/Perry/internal/adapter/validator"
	"github.com/wang900115/Perry/internal/domain/entity"
	gorminterface "github.com/wang900115/Perry/internal/domain/interface/gorm"
	redisinterface "github.com/wang900115/Perry/internal/domain/interface/redis"
)

type ToDo struct {
	// db
	todoRepo gorminterface.ToDo
	// redis
	todoCache redisinterface.ToDo
}

func NewToDoUsecase(todoRepo *gorminterface.ToDo, todoCache *redisinterface.ToDo) *ToDo {
	return &ToDo{todoRepo: *todoRepo, todoCache: *todoCache}
}

// 新增待辦任務
func (t *ToDo) Create(ctx context.Context, user_id uint, input validator.ToDoCreateRequest) (*entity.ToDo, error) {
	todo, err := t.todoRepo.Create(ctx, user_id, input)
	if err != nil {
		return nil, err
	}
	err = t.todoCache.Initialize(ctx, todo.Id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

// 刪除待辦任務
func (t *ToDo) Delete(ctx context.Context, input validator.ToDoDeleteRequest) error {
	err := t.todoCache.Delete(ctx, input)
	if err != nil {
		return err
	}
	return t.todoRepo.Delete(ctx, input)
}

// 更新待辦任務
func (t *ToDo) Update(ctx context.Context, input validator.ToDoUpdateRequest) (*entity.ToDo, error) {
	if err := t.todoCache.Initialize(ctx, input.ID); err != nil {
		return nil, err
	}
	return t.todoRepo.Update(ctx, input)
}

// 取得該使用者的待辦任務
func (t *ToDo) Query(ctx context.Context, user_id uint) ([]*entity.ToDo, error) {
	return t.todoRepo.Query(ctx, user_id)
}

// 取得任務狀態
func (t *ToDo) GetStatus(ctx context.Context, input validator.ToDoGetStatusRequest) (string, error) {
	status, err := t.todoCache.GetStatus(ctx, input)
	if err != nil {
		return "", err
	}
	if len(status) > 0 {
		return status, nil
	}
	return t.todoRepo.GetStatus(ctx, input)
}
