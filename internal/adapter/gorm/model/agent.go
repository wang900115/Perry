package gormmodel

import (
	"github.com/wang900115/Perry/internal/domain/entity"
	"gorm.io/gorm"
)

// #region 表格欄位
type Agent struct {
	gorm.Model

	UserID uint `gorm:"column:user_id"`
	User   User `gorm:"foreignKey:UserID"`

	Name        string `gorm:"column:name"`
	Age         int    `gorm:"column:age"`
	Role        string `gorm:"column:role"`
	Language    string `gorm:"column:language"`
	Description string `gorm:"column:description"`

	Status string `gorm:"column:status"`
}

// #endregion

// #region 表格名稱
func (Agent) TableName() string {
	return "agent"
}

// #endregion

// #region -> Domain
func (a *Agent) Domain() *entity.Agent {
	return &entity.Agent{
		ID:          a.ID,
		Name:        a.Name,
		Age:         a.Age,
		Role:        a.Role,
		Language:    a.Language,
		Description: a.Description,

		Status: a.Status,
	}
}
