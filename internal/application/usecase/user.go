package usecase

import (
	"context"
	"time"

	"github.com/wang900115/Perry/internal/adapter/validator"
	"github.com/wang900115/Perry/internal/domain/entity"
	gorminterface "github.com/wang900115/Perry/internal/domain/interface/gorm"
	redisinterface "github.com/wang900115/Perry/internal/domain/interface/redis"
)

type User struct {
	// db
	userRepo gorminterface.User
	// redis
	tokenRepo redisinterface.Token
	// redis
	sessionRepo redisinterface.Session
}

func NewUserUsecase(userRepo gorminterface.User, tokenRepo redisinterface.Token, sessionRepo redisinterface.Session) *User {
	return &User{userRepo: userRepo, tokenRepo: tokenRepo, sessionRepo: sessionRepo}
}

// 註冊
func (u *User) Register(ctx context.Context, input validator.RegisterRequest) error {
	err := u.userRepo.Register(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

// 登入 (單點)
func (u *User) Login(ctx context.Context, username, password, ip, userAgent string) (*entity.UserStatus, string, error) {
	userSatatus, err := u.userRepo.Login(ctx, username, password, ip, time.Now())
	if err != nil {
		return nil, "", err
	}
	userId := userSatatus.UserId
	sessionId, err := u.sessionRepo.Generate(ctx, userId, ip, userAgent)
	if err != nil {
		return nil, "", err
	}
	token, err := u.tokenRepo.Generate(ctx, userId, sessionId)
	if err != nil {
		_ = u.sessionRepo.Delete(ctx, sessionId)
		return nil, "", err
	}
	return userSatatus, token, nil
}

// 登出 (單點)
func (u *User) Logout(ctx context.Context, user_id uint, session_id int64) error {
	err := u.sessionRepo.Deactivate(ctx, session_id)
	if err != nil {
		return err
	}
	err = u.tokenRepo.Delete(ctx, user_id, session_id)
	if err != nil {
		return err
	}
	err = u.userRepo.UpdateLastLogout(ctx, user_id, time.Now())
	if err != nil {
		return err
	}
	return nil
}

// 刪除帳號
func (u *User) Delete(ctx context.Context, user_id uint) error {
	err := u.userRepo.Delete(ctx, user_id)
	if err != nil {
		return err
	}
	err = u.tokenRepo.DeleteAll(ctx, user_id)
	if err != nil {
		return err
	}
	return nil
}

// 更新設定
func (u *User) UpdateSettins(ctx context.Context, user_id uint, input validator.UpdateSettingsRequest) (*entity.User, error) {
	user, err := u.userRepo.UpdateSettings(ctx, user_id, input)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// 更新密碼
func (u *User) UpdatePassword(ctx context.Context, user_id uint, newPassword string) error {
	err := u.userRepo.UpdatePassword(ctx, user_id, newPassword)
	if err != nil {
		return err
	}
	return nil
}

// !todo 找回帳號
func (u *User) Forgot(ctx context.Context) error {
	return nil
}
