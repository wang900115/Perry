package gormimplement

import (
	"context"
	"errors"
	"time"

	gormmodel "github.com/wang900115/Perry/internal/adapter/gorm/model"
	"github.com/wang900115/Perry/internal/adapter/validator"
	"github.com/wang900115/Perry/internal/domain/entity"
	gorminterface "github.com/wang900115/Perry/internal/domain/interface/gorm"
	"github.com/wang900115/utils/encrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm *gorm.DB
}

func NewUserImplement(gorm *gorm.DB) gorminterface.User {
	return &User{gorm: gorm}
}

func (u *User) Register(ctx context.Context, input validator.RegisterRequest) error {
	password, err := encrypt.HashPasswordArgon2id(input.Password)
	if err != nil {
		return err
	}
	createdUser := gormmodel.User{
		Username:  input.Username,
		Password:  password,
		FullName:  input.FullName,
		NickName:  input.NickName,
		AvatarURL: input.AvatarURL,
		Phone:     input.Phone,
		Email:     input.Email,
		Location: gormmodel.Location{
			Country:   input.Location.Country,
			City:      input.Location.City,
			Latitude:  input.Location.Latitude,
			Longitude: input.Location.Longitude,
		},
		Description: input.Description,
	}
	if err := u.gorm.WithContext(ctx).Create(createdUser).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) Login(ctx context.Context, username, password, ip string, lastLogin time.Time) (*entity.UserStatus, error) {
	var userStatusModel gormmodel.UserStatus
	if err := u.gorm.WithContext(ctx).Where("username = ?", username).First(&userStatusModel).Error; err != nil {
		return nil, err
	}
	right, err := encrypt.VerifyPasswordArgon2id(userStatusModel.Password, password)
	if err != nil {
		return nil, err
	}
	if !right {
		return nil, errors.New("password incorrect")
	}
	if err := u.gorm.WithContext(ctx).Model(&gormmodel.UserStatus{}).Where("username = ?", username).Updates(map[string]interface{}{
		"last_ip":    ip,
		"last_login": lastLogin,
	}).Error; err != nil {
		return nil, err
	}
	return userStatusModel.ToDomain(), nil
}

func (u *User) UpdateLastLogout(ctx context.Context, user_id uint, lastLogout time.Time) error {
	if err := u.gorm.WithContext(ctx).Table("user").Where("user_id = ?", user_id).Update("last_logout", lastLogout).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) Delete(ctx context.Context, user_id uint) error {
	var userStatusModel gormmodel.UserStatus
	if err := u.gorm.WithContext(ctx).Delete(&userStatusModel, user_id).Error; err != nil {
		return err
	}
	return nil
}
