package usecase

import (
	"context"

	"github.com/wang900115/Perry/internal/adapter/validator"
	"github.com/wang900115/Perry/internal/domain/entity"
	gorminterface "github.com/wang900115/Perry/internal/domain/interface/gorm"
	redisinterface "github.com/wang900115/Perry/internal/domain/interface/redis"
)

type ToDo struct {
	todoRepo  gorminterface.ToDo
	todoCache redisinterface.ToDo
}

func NewToDoUsecase(todoRepo *gorminterface.ToDo, todoCache *redisinterface.ToDo) *ToDo {
	return &ToDo{todoRepo: *todoRepo, todoCache: *todoCache}
}

// 新增待辦任務(先在db新增再redis初始化)
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

// 刪除待辦任務(先刪db再刪redis)
func (t *ToDo) Delete(ctx context.Context, input validator.ToDoDeleteRequest) error {
	err := t.todoRepo.Delete(ctx, input)
	if err != nil {
		return err
	}
	return t.todoCache.Delete(ctx, input)
}

// 更新待辦任務(先db更新狀態 再redis更新todo資訊)
func (t *ToDo) Update(ctx context.Context, input validator.ToDoUpdateRequest) (*entity.ToDo, error) {
	_, err := t.todoRepo.Update(ctx, input)
	if err != nil {
		return nil, err
	}
	todoModel, err := t.todoCache.Update(ctx, input)
	if err != nil {
		return nil, err
	}
	return todoModel.ToDomain(), nil
}

// 取得該使用者的所有待辦任務
func (t *ToDo) Query(ctx context.Context, user_id uint) ([]*entity.ToDo, error) {
	todoModels, err := t.todoCache.Get(ctx, user_id)
	if err != nil {
		return nil, err
	}
	if len(todoModels) > 0 {
		todos := make([]*entity.ToDo, 0, len(todoModels))
		for _, todoModel := range todoModels {
			todos = append(todos, todoModel.ToDomain())
		}
		return todos, nil
	}
	return t.todoRepo.Query(ctx, user_id)
}

// 取得該代理任務資訊 (先去redis查找如果有則回傳，如果沒有則去DB查找)
func (t *ToDo) Get(ctx context.Context, agent_id uint) ([]*entity.ToDo, error) {
	todoModels, err := t.todoCache.Get(ctx, agent_id)
	if err != nil {
		return nil, err
	}
	if len(todoModels) > 0 {
		todos := make([]*entity.ToDo, 0, len(todoModels))
		for _, todoModel := range todoModels {
			todos = append(todos, todoModel.ToDomain())
		}
		return todos, nil
	}
	return t.todoRepo.Query(ctx, agent_id)
}
