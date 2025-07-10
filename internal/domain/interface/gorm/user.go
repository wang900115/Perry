package gorminterface

import (
	"context"
	"time"

	"github.com/wang900115/Perry/internal/adapter/validator"
	"github.com/wang900115/Perry/internal/domain/entity"
)

type User interface {
	//  註冊
	Register(context.Context, validator.RegisterRequest) error
	//  登入(觸發更新狀態)
	Login(context.Context, string, string, string, time.Time) (*entity.UserStatus, error)
	//  登出(觸發更新狀態)
	UpdateLastLogout(context.Context, uint, time.Time) error
	//  刪除帳號(依據id)
	Delete(context.Context, uint) error
	// 更新設定
	UpdateSettings(context.Context, uint, validator.UpdateSettingsRequest) (*entity.User, error)
	// 更新密碼
	UpdatePassword(context.Context, uint, string) error
}
