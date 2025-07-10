package gormmodel

import (
	"time"

	"github.com/wang900115/Perry/internal/domain/entity"
	"gorm.io/gorm"
)

// #region 表格欄位
type ToDo struct {
	gorm.Model

	UserID uint `gorm:"column:user_id"`
	User   User `gorm:"foreignKey:UserID"`

	Name      string    `gorm:"column:name"`
	Priority  string    `gorm:"priority"`
	StartTime time.Time `gorm:"column:start_time"`
	EndTime   time.Time `gorm:"column:end_time"`
}

// #endregion

// #region 表格名稱
func (ToDo) TableName() string {
	return "user_todo"
}

// #endregion

// #region -> Domain

func (t *ToDo) ToDomain() *entity.ToDo {
	return &entity.ToDo{
		Id:        t.ID,
		Name:      t.Name,
		Priority:  t.Priority,
		StartTime: t.StartTime,
		EndTime:   t.EndTime,
	}
}

// #endregion
