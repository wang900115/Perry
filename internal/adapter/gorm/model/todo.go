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

	AgentID uint  `gorm:"column:agent_id"`
	Agent   Agent `gorm:"foreignKey:AgentID"`

	Name      string    `gorm:"column:name"`
	Priority  string    `gorm:"column:priority"`
	StartTime time.Time `gorm:"column:start_time"`
	EndTime   time.Time `gorm:"column:end_time"`
	Status    string    `gorm:"column:status"`
}

// #endregion

// #region 表格名稱
func (ToDo) TableName() string {
	return "todo"
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
		Status:    t.Status,
	}
}

// #endregion
