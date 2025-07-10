package usecase

import (
	"context"

	"github.com/wang900115/Perry/internal/adapter/validator"
	"github.com/wang900115/Perry/internal/domain/entity"
	gorminterface "github.com/wang900115/Perry/internal/domain/interface/gorm"
)

type ToDo struct {
	// db
	todoRepo gorminterface.ToDo
}

func NewToDoUsecase(todoRepo gorminterface.ToDo) *ToDo {
	return &ToDo{todoRepo: todoRepo}
}

// 新增待辦事項
func (t *ToDo) Create(ctx context.Context, user_id uint, input validator.ToDoCreateRequest) (*entity.ToDo, error) {
	return t.todoRepo.Create(ctx, user_id, input)
}

// 刪除待辦事項
func (t *ToDo) Delete(ctx context.Context, user_id uint, input validator.ToDoDeleteRequest) error {
	return t.todoRepo.Delete(ctx, user_id, input)
}

// 更新待辦事項
func (t *ToDo) Update(ctx context.Context, user_id uint, input validator.ToDoUpdateRequest) (*entity.ToDo, error) {
	return t.todoRepo.Update(ctx, user_id, input)
}

// 取得該使用者的待辦事項
func (t *ToDo) Query(ctx context.Context, user_id uint) ([]*entity.ToDo, error) {
	return t.todoRepo.Query(ctx, user_id)
}
